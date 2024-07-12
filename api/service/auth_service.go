package service

import (
	"api_go/api/model"
	"api_go/api/repository"
	"api_go/api/request"
	"api_go/pkg/db"
	"errors"
	"time"
)

func RegisterUser(request request.RegisterRequest) error {
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if request.Password != request.PasswordConfirmation {
		return errors.New("两次输入的密码不一致")
	}
	user := model.User{
		Mobile:    request.Mobile,
		Password:  request.Password,
		CreatedAt: time.Now(),
	}
	error := db.DB.Select("Mobile", "Password", "CreatedAt").Create(&user).Error

	if error != nil {
		return errors.New("注册失败")
	}

	member := model.Member{
		UserId:    user.ID,
		Mobile:    request.Mobile,
		CreatedAt: time.Now(),
	}
	error2 := db.DB.Select("UserId", "Mobile", "CreatedAt").Create(&member).Error
	if error2 != nil {
		return error2
	}

	return nil
}

func LoginUser(mobile, password string) (*model.User, error) {
	user, err := repository.GetUserByMobile(mobile)
	if err != nil {
		return nil, errors.New("不存在该用户")
	}
	/*
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return nil, errors.New("invalid password")
		}
	*/

	if user.Password != password {
		return nil, errors.New("不存在用户或密码不正确")
	}

	return user, nil
}
