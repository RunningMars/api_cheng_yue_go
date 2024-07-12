package controller

import (
	"api_go/api/util"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "未上传图片"})
		return
	}
	defer file.Close()

	// Create a unique filename
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), header.Filename)
	filename = "image/" + filename

	// Get OSS bucket
	bucket, err := util.GetBucket()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OSS初始化失败"})
		return
	}

	// Upload file to OSS
	err = bucket.PutObject(filename, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "上传图片失败"})
		return
	}

	//fileURL := fmt.Sprintf("https://%s.%s/%s", config.OSS.BucketName, config.OSS.Endpoint, filename)
	url := "https://cheng-yue.oss-cn-chengdu.aliyuncs.com/" + filename

	var maps = make(map[string]interface{})
	maps["url"] = url
	maps["folder"] = "image"
	c.JSON(http.StatusOK, util.Success(maps))
}
