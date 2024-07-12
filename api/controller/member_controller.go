package controller

import (
	"api_go/api/model"
	"api_go/api/request"
	"api_go/api/service"
	"api_go/api/util"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MemberList(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	var memberListReqeust request.MemberListReqeust
	error := c.ShouldBindQuery(&memberListReqeust)
	if error != nil {
		c.JSON(http.StatusBadRequest, util.Abort(error.Error()))
		return
	}

	sex, exists := c.Get("sex")
	fmt.Println("sex", sex)
	if exists {
		if sex.(int) == 1 {
			memberListReqeust.Sex = 2
		} else if sex.(int) == 2 {
			memberListReqeust.Sex = 1
			// } else {
			// 	c.JSON(http.StatusBadRequest, gin.H{
			// 		"code":    1,
			// 		"message": "请先设置性别",
			// 	})
			// 	return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "请先设置性别",
			"error":   err.Error(),
		})
		return
	}

	memberId, exists := c.Get("memberId")
	if exists {
		memberListReqeust.MyMemberId = memberId.(uint)
	}
	fmt.Println("memberListReqeust", memberListReqeust)
	fmt.Println("memberId", memberId)
	members, total, err := service.GetMembers(pageNum, pageSize, memberListReqeust)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Abort(err.Error()))
		return
	}

	var pageInfo map[string]any
	pageInfo = make(map[string]any)
	pageInfo["total"] = total
	pageInfo["records"] = members
	pageInfo["size"] = pageSize
	pageInfo["current"] = pageNum

	c.JSON(http.StatusOK, util.Success(pageInfo))
	return
}

func MemberDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, util.Abort("id 格式不正确"))
		return
	}
	memberId, exists := c.Get("memberId")
	if !exists {
		c.JSON(http.StatusOK, util.Abort("未发现当前会员member_id信息"))
		return
	}
	member, error := service.GetMemberById(uint(id), memberId.(uint))
	if error != nil {
		c.JSON(http.StatusOK, util.Abort("获取用户信息失败"))
		return
	}
	c.JSON(http.StatusOK, util.Success(member))
	return
}

func MemberSave(c *gin.Context) {
	var member model.Member
	error := c.ShouldBindJSON(&member)
	if error != nil {
		c.JSON(http.StatusBadRequest, util.Abort(error.Error()))
		return
	}
	memberId, exists := c.Get("memberId")
	if !exists {
		c.JSON(http.StatusOK, util.Abort("未发现当前会员member_id信息"))
		return
	}
	if memberId.(uint) != member.ID {
		c.JSON(http.StatusOK, util.Abort("没有修改权限"))
		return
	}
	//fmt.Println("member", member)
	error = service.SaveMember(member)
	if error != nil {
		c.JSON(http.StatusInternalServerError, util.Abort(error.Error()))
		return
	}
	c.JSON(http.StatusOK, util.Success(nil))
	return
}

func ThumbsUpList(c *gin.Context) {
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

	list, total, error2 := service.GetThumbsUpList(pageNum, pageSize, memberId.(uint), keyWord)
	if error2 != nil {
		c.JSON(http.StatusOK, util.Failed("查询失败,请重试"))
		return
	}

	pageInfo := util.NewPage(list, pageNum, pageSize, total)
	var res map[string]any
	res = make(map[string]any)
	res["thumbsUpListPage"] = pageInfo
	res["allThumbsUpCount"] = total
	c.JSON(http.StatusOK, util.Success(res))
}

func UpdateFavorite(c *gin.Context) {
	isFavorite := c.Query("isFavorite")

	toMemberId, err := strconv.Atoi(c.Query("toMemberId"))
	if err != nil || toMemberId <= 0 {
		c.JSON(http.StatusOK, util.Abort("请求数据格式错误"))
		return
	}

	var updateFavorite int
	if isFavorite == "1" {
		updateFavorite = 1
	} else if isFavorite == "0" {
		updateFavorite = 0
	} else {
		c.JSON(http.StatusOK, util.Abort("请求数据格式错误."))
		return
	}

	memberId, exists := c.Get("memberId")
	if !exists {
		c.JSON(http.StatusOK, util.Abort("未发现当前会员member_id信息"))
		return
	}
	error := service.UpdateFavorite(memberId.(uint), uint(toMemberId), updateFavorite)
	if error != nil {
		c.JSON(http.StatusOK, util.Abort(error.Error()))
		return
	}
	c.JSON(http.StatusOK, util.Success(nil))
}

func UpdateThumbsUp(c *gin.Context) {
	isThumbsUp := c.Query("isThumbsUp")

	toMemberId, err := strconv.Atoi(c.Query("toMemberId"))
	if err != nil || toMemberId <= 0 {
		c.JSON(http.StatusOK, util.Abort("请求数据格式错误"))
		return
	}

	var updateThumbsUp int
	if isThumbsUp == "1" {
		updateThumbsUp = 1
	} else if isThumbsUp == "0" {
		updateThumbsUp = 0
	} else {
		c.JSON(http.StatusOK, util.Abort("请求数据格式错误."))
		return
	}

	memberId, exists := c.Get("memberId")
	if !exists {
		c.JSON(http.StatusOK, util.Abort("未发现当前会员member_id信息"))
		return
	}
	error := service.UpdateThumbsUp(memberId.(uint), uint(toMemberId), updateThumbsUp)
	if error != nil {
		c.JSON(http.StatusOK, util.Abort(error.Error()))
		return
	}
	c.JSON(http.StatusOK, util.Success(nil))
}

func Test(c *gin.Context) {
	var memberListReqeust request.MemberListReqeust
	error := c.ShouldBindQuery(&memberListReqeust)
	if error != nil {
		c.JSON(http.StatusBadRequest, util.Abort(error.Error()))
		return
	}

	fmt.Println("memberListReqeust", memberListReqeust)

	age := c.Query("age")              // age参数为int类型
	price := c.Query("price")          // price参数为float64类型
	name := c.Query("name")            // name参数为string类型
	details := c.QueryArray("details") // details参数为[]string类型
	fmt.Printf("age type:%T\n", age)
	fmt.Println("age", age)
	fmt.Printf("price type:%T\n", price)
	fmt.Println("price", price)
	fmt.Printf("name type: %T\n", name)
	fmt.Println("name", name)
	fmt.Printf("details type:%T\n", details)
	fmt.Println("details", details)

	os.Exit(2)
}
