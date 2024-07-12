package controller

import (
	"api_go/api/request"
	"api_go/api/service"
	"api_go/api/util"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ChatList(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	keyWord := c.Query("keyWord")

	memberId, exists := c.Get("memberId")
	if !exists {
		c.JSON(http.StatusOK, util.Abort("未发现当前会员member_id信息"))
		return
	}

	var chatRooms, total, error = service.GetChatList(memberId.(uint), pageNum, pageSize, keyWord)

	if error != nil {
		c.JSON(http.StatusOK, util.Failed("查询失败,请重试"))
		return
	}

	pageInfo := util.NewPage(chatRooms, pageNum, pageSize, total)
	c.JSON(http.StatusOK, util.Success(pageInfo))
}

func MessageList(c *gin.Context) {
	chatRoomId, err := strconv.Atoi(c.Query("chatRoomId"))
	toMemberId, err := strconv.Atoi(c.Query("toMemberId"))
	if chatRoomId == 0 && toMemberId == 0 {
		c.JSON(http.StatusOK, util.Abort("请求参数格式不正确"))
	}

	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	memberId, exists := c.Get("memberId")
	if !exists {
		c.JSON(http.StatusOK, util.Abort("未发现当前会员member_id信息"))
		return
	}

	map1 := make(map[string]interface{})
	map1["chatRoomId"] = chatRoomId
	map1["toMemberId"] = toMemberId
	map1["pageNum"] = pageNum
	map1["pageSize"] = pageSize
	map1["memberId"] = memberId.(uint)

	var chatRoomMessages, error = service.GetChatRoomMessageList(map1)

	if error != nil {
		c.JSON(http.StatusOK, util.Failed("查询失败,请重试"))
		return
	}

	c.JSON(http.StatusOK, util.Success(chatRoomMessages))
}

func SendMessage(c *gin.Context) {
	var sendMessageRequest request.SendMessage
	c.ShouldBindJSON(&sendMessageRequest)

	memberId, exists := c.Get("memberId")
	if !exists {
		c.JSON(http.StatusOK, util.Abort("未发现当前会员member_id信息"))
		return
	}
	sendMessageRequest.MemberId = memberId.(uint)

	var error = service.SendMessage(sendMessageRequest)

	if error != nil {
		c.JSON(http.StatusOK, util.Failed("发送失败,请重试"))
		return
	}

	c.JSON(http.StatusOK, util.Success(nil))
}

func UnreadCount(c *gin.Context) {
	memberId, exists := c.Get("memberId")
	if !exists {
		c.JSON(http.StatusOK, util.Abort("未发现当前会员member_id信息"))
		return
	}

	var count, error = service.GetUnreadCount(memberId.(uint))

	if error != nil {
		c.JSON(http.StatusOK, util.Failed("获取失败,请重试"))
		return
	}
	maps := make(map[string]interface{})
	maps["unreadChatCount"] = count
	c.JSON(http.StatusOK, util.Success(maps))
}

func ReadAll(c *gin.Context) {
	memberId, exists := c.Get("memberId")
	if !exists {
		c.JSON(http.StatusOK, util.Abort("未发现当前会员member_id信息"))
		return
	}

	var error = service.ReadAll(memberId.(uint))

	if error != nil {
		c.JSON(http.StatusOK, util.Failed("操作失败,请重试"))
		return
	}

	c.JSON(http.StatusOK, util.Success(nil))
}
