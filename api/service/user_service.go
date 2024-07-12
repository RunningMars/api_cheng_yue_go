package service

import (
	"api_go/api/model"
	"api_go/api/repository"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetAllUsers() []model.User {
	// 假设这是从数据库获取数据
	return []model.User{
		{ID: 1, Name: "John 1"},
		{ID: 2, Name: "Jane 2"},
	}
}

func GetUsers(page, pageSize int, keyword string) ([]model.User, error) {
	return repository.GetUsers(page, pageSize, keyword)
}

func CreateUser(user User) {
	// 假设这是保存数据到数据库
	//model.SaveUser(user)
}
