package main

import (
	"api_go/api/router"
	"api_go/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库连接
	db.Init()
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	router.SetupRouter(r)
	r.Run(":8090") // 监听并在 0.0.0.0:8060 上启动服务

}
