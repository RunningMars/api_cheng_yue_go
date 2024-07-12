package service

import (
	"api_go/api/model"
	"api_go/config"
	"api_go/pkg/db"
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/jinzhu/gorm"
)

const CodeLength = 4
const ExpiryDuration = 10 * time.Minute
const MaxSendPerHour = 20

func generateRandomCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < length; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}

func SendSMS(mobile string) error {
	now := time.Now()

	var exists int64
	if err := db.DB.Model(&model.MobileValidateCode{}).Where("mobile = ?", mobile).
		Where("created_at > ?", now.Add(-1*time.Minute)).Count(&exists).Error; err != nil {
		return err
	}
	if exists > 0 {
		return errors.New("请稍后后重试,重新发送短信验证码.")
	}

	//startOfHour := now.Truncate(time.Hour)

	// var count int64
	// if err := db.DB.Model(&model.SendSMS{}).
	// 	Where("mobile = ? AND created_at >= ?", mobile, startOfHour).
	// 	Count(&count).Error; err != nil {
	// 	return err
	// }

	// if count >= MaxSendPerHour {
	// 	return fmt.Errorf("exceeded max send limit for the hour")
	// }

	code := generateRandomCode(CodeLength)

	// Initialize SMS client
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", config.AliyunSMS.AccessKeyID, config.AliyunSMS.AccessKeySecret)
	if err != nil {
		return err
	}

	templateParam := map[string]string{
		"code": code,
	}
	templateParamStr, _ := json.Marshal(templateParam)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = mobile
	request.SignName = config.AliyunSMS.SignName
	request.TemplateCode = config.AliyunSMS.TemplateCode
	request.TemplateParam = string(templateParamStr)

	response, err := client.SendSms(request)
	if err != nil {
		return err
	}

	// Save to database
	sendSMS := model.SendSMS{
		Mobile:        mobile,
		TemplateCode:  config.AliyunSMS.TemplateCode,
		SignName:      config.AliyunSMS.SignName,
		TemplateParam: string(templateParamStr),
		Message:       response.Message,
		Response:      response.GetHttpContentString(),
		Code:          code,
		BizID:         response.BizId,
		RequestID:     response.RequestId,
		Status:        response.Code,
		CreatedAt:     now,
	}

	if err := db.DB.Create(&sendSMS).Error; err != nil {
		return err
	}

	// Save validation code
	validateCode := model.MobileValidateCode{
		Mobile:    mobile,
		Code:      code,
		CreatedAt: now,
	}

	if err := db.DB.Create(&validateCode).Error; err != nil {
		return err
	}

	return nil
}

func ValidateSMSCode(mobile, code string) (bool, error) {
	var validateCode model.MobileValidateCode
	if err := db.DB.Where("mobile = ? AND code = ?", mobile, code).Order("created_at DESC").First(&validateCode).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	if time.Since(validateCode.CreatedAt) > ExpiryDuration {
		return false, nil
	}

	return true, nil
}
