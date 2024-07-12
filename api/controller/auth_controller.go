package controller

import (
	"api_go/api/model"
	"api_go/api/request"
	"api_go/api/service"
	"api_go/api/util"
	"api_go/pkg/db"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("r909jfawfh983wf9")

type Claims struct {
	UserId   uint   `json:"userId"`
	Name     string `json:"name"`
	MemberId uint   `json:"memberId"`
	Sex      int    `json:"sex"`

	jwt.RegisteredClaims
}

func Register(c *gin.Context) {
	var request request.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, util.Abort(err.Error()))
		return
	}

	var exists int64
	if err := db.DB.Model(&model.User{}).Where("mobile = ?", request.Mobile).
		Count(&exists).Error; err != nil {
		c.JSON(http.StatusOK, util.Abort(err.Error()))
	}
	if exists > 0 {
		c.JSON(http.StatusOK, util.Abort("此手机号码已注册,请直接登录."))
		return
	}

	//验证 otp
	valid, err := service.ValidateSMSCode(request.Mobile, request.Code)
	if err != nil {
		c.JSON(http.StatusOK, util.Abort("验证码不正确"))
		return
	}

	if valid == false {
		c.JSON(http.StatusOK, util.Abort("验证码未通过验证"))
		return
	}
	if err := service.RegisterUser(request); err != nil {
		c.JSON(http.StatusOK, util.Abort(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.Success("注册成功"))
}

func Login(c *gin.Context) {
	var request struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, util.Abort(err.Error()))
		return
	}

	user, err := service.LoginUser(request.Mobile, request.Password)
	if err != nil {
		// c.JSON(http.StatusUnauthorized, util.Abort(err.Error()))
		c.JSON(http.StatusOK, util.Abort(err.Error()))
		return
	}

	member, err := service.GetMemberByUserId(user.ID)
	if err != nil {
		c.JSON(http.StatusOK, util.Abort(err.Error()))
		return
	}
	//30 days
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &Claims{
		UserId:   user.ID,
		Name:     user.Mobile,
		MemberId: member.ID,
		Sex:      member.Sex,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Abort(err.Error()))
		return
	}

	var data map[string]any
	data = make(map[string]any) // 使用make函数初始化map
	data["userInfo"] = map[string]any{
		"id":     user.ID,
		"mobile": user.Mobile,
		"name":   user.Name,
		"member": member,
	}
	data["accessToken"] = tokenString
	data["tokenType"] = "bearer"

	c.JSON(http.StatusOK, util.Success(data))
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, util.Success(nil))
}

func SendOtp(c *gin.Context) {
	var maps map[string]any
	c.ShouldBindJSON(&maps)

	var exists int64
	if err := db.DB.Model(&model.User{}).Where("mobile = ?", maps["mobile"].(string)).
		Count(&exists).Error; err != nil {
		c.JSON(http.StatusOK, util.Abort(err.Error()))
	}
	if exists > 0 {
		c.JSON(http.StatusOK, util.Abort("此手机号码已注册,请直接登录."))
		return
	}

	error := service.SendSMS(maps["mobile"].(string))
	if error != nil {
		c.JSON(http.StatusOK, util.Abort(error.Error()))
		return
	}

	c.JSON(http.StatusOK, util.Success(nil))
}
