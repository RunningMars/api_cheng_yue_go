package controller

import (
	"api_go/api/model"
	"api_go/api/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	keyword := c.Query("keyword")

	users, err := service.GetUsers(page, pageSize, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}


func GetUserTest(c *gin.Context) {
	fmt.Println("Index...")
	var userList []model.User
	userList = service.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"message": "welcome Index.", "data": userList})
}

