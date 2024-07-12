package repository

import (
	"api_go/api/model"
	"api_go/pkg/db"
)

func GetUsers(page, pageSize int, keyword string) ([]model.User, error) {
	var users []model.User
	offset := (page - 1) * pageSize
	query := db.DB.Offset(offset).Limit(pageSize)

	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByMobile(mobile string) (*model.User, error) {
	var user model.User
	if err := db.DB.Where("mobile = ?", mobile).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *model.User) error {
	return db.DB.Create(user).Error
}
