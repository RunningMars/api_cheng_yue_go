package service

import (
	"api_go/api/model"
	"api_go/api/repository"
	"api_go/api/request"
)

func GetMemberByUserId(userId uint) (*model.Member, error) {
	return repository.GetMemberByUserId(userId)
}

func GetMembers(page, pageSize int, memberListReqeust request.MemberListReqeust) (*[]model.Member, int64, error) {
	return repository.GetMembers(page, pageSize, memberListReqeust)
}

func GetMemberById(memberId uint, currentMemberId uint) (*model.Member, error) {
	return repository.GetMemberById(memberId, currentMemberId)
}

func SaveMember(member model.Member) error {
	return repository.SaveMember(member)
}

func GetThumbsUpList(pageNum, pageSize int, memberId uint, keyWord string) (*[]map[string]any, int64, error) {
	return repository.GetThumbsUpList(pageNum, pageSize, memberId, keyWord)
}

func UpdateThumbsUp(memberId, toMemberId uint, updateThumbsUp int) error {
	return repository.UpdateThumbsUp(memberId, toMemberId, updateThumbsUp)
}

func UpdateFavorite(memberId, toMemberId uint, updateFavorite int) error {
	return repository.UpdateFavorite(memberId, toMemberId, updateFavorite)
}
