package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	fmt.Println("Index...")

	fmt.Println(c.Get("userId"))
	fmt.Println(c.Get("name"))
	fmt.Println(c.Get("memberId"))
	fmt.Println(c.Get("sex"))

	c.JSON(http.StatusOK, gin.H{"message": "welcome Index."})
}

func Test2(c *gin.Context) {
	fmt.Println("Test...")
	c.JSON(http.StatusOK, gin.H{"message": "welcome Test."})
}
