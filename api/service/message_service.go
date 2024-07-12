package service

import (
	"api_go/api/model"
	"api_go/api/repository"
	"api_go/api/request"
	"api_go/api/util"
	"api_go/pkg/db"
	"fmt"

	// "fmt"
	"strconv"
	"time"
)

func GetChatList(memberId uint, pageNum, pageSize int, keyWord string) (*[]model.ChatRoom, int64, error) {
	return repository.GetChatList(memberId, pageNum, pageSize, keyWord)
}

func GetChatRoomMessageList(maps map[string]any) (map[string]any, error) {

	chatRoomId := maps["chatRoomId"].(int)
	toMemberId := maps["toMemberId"].(int)
	memberId := maps["memberId"].(uint)

	if chatRoomId == 0 {
		//查找是否已存在 chatRoom
		existsChatRoom, error := repository.FindChatRoomByMemberIds(memberId, uint(toMemberId))
		if error != nil {
			return nil, error
		}
		//如何不存在则返回空数据
		if existsChatRoom.ID == 0 {
			member, _ := repository.GetMemberById2(uint(toMemberId))
			result := make(map[string]any)
			result["data"] = make([]any, 0)
			result["toMember"] = map[string]any{
				"id":           member.ID,
				"nickName":     member.NickName,
				"profilePhoto": member.ProfilePhoto,
			}
			return result, nil
		} else {
			maps["chatRoomId"] = existsChatRoom.ID
		}
	} else {
		maps["chatRoomId"] = uint(chatRoomId)
	}

	chatRoomMessages, total, error := repository.GetChatRoomMessageList(maps)
	if error != nil {
		return nil, error
	}

	//将当天聊天设为已读
	repository.ReadChatRoom(memberId, maps["chatRoomId"].(uint))

	var oppsiteMemberId uint
	var oppsiteMember map[string]any = make(map[string]any)
	memberIds := make([]uint, 0)
	for _, chatRoomMessage := range *chatRoomMessages {
		memberIds = append(memberIds, chatRoomMessage.ToMemberId)
		memberIds = append(memberIds, chatRoomMessage.FromMemberId)
		if chatRoomMessage.ToMemberId != memberId {
			oppsiteMemberId = chatRoomMessage.ToMemberId
		}
	}

	members, _ := repository.GetMemberByIds(memberIds)
	membersMap := make(map[uint]model.Member, len(*members))
	for _, member := range *members {
		membersMap[member.ID] = member
	}

	for k, chatRoomMessage := range *chatRoomMessages {
		if member, found := membersMap[chatRoomMessage.FromMemberId]; found {
			(*chatRoomMessages)[k].FromMember = map[string]any{
				"id":           member.ID,
				"nickName":     member.NickName,
				"profilePhoto": member.ProfilePhoto,
			}
		}
		if member, found := membersMap[chatRoomMessage.ToMemberId]; found {
			(*chatRoomMessages)[k].ToMember = map[string]any{
				"id":           member.ID,
				"nickName":     member.NickName,
				"profilePhoto": member.ProfilePhoto,
			}
		}
	}
	if member, found := membersMap[oppsiteMemberId]; found {
		oppsiteMember = map[string]any{
			"id":           member.ID,
			"nickName":     member.NickName,
			"profilePhoto": member.ProfilePhoto,
		}
	}

	result := make(map[string]any)
	result["data"] = util.NewPage[model.ChatRoomMessage](chatRoomMessages, maps["pageNum"].(int), maps["pageSize"].(int), total)
	result["toMember"] = oppsiteMember
	return result, nil
}

func SendMessage(sendMessage request.SendMessage) error {
	toMemberId, _ := strconv.ParseUint(sendMessage.ToMemberId, 10, 32) // 将字符串转换为uint
	//查找是否已存在 chatRoom
	existsChatRoom, error := repository.FindChatRoomByMemberIds(sendMessage.MemberId, uint(toMemberId))
	if error != nil {
		return error
	}
	var chatRoomId uint
	//没有则创建 chatRoom
	if existsChatRoom.ID == 0 {
		var str string
		str = fmt.Sprintf("%s_%s", strconv.FormatUint(uint64(sendMessage.MemberId), 10), sendMessage.ToMemberId)

		var chatRoom = model.ChatRoom{
			ChatRoomName: str,
			CreatedAt:    time.Now(),
		}
		if err := db.DB.Select("ChatRoomName", "CreatedAt").Create(&chatRoom).Error; err != nil {
			return err
		}
		var chatRoomMemberMe = model.ChatRoomMember{
			ChatRoomId: chatRoom.ID,
			MemberId:   sendMessage.MemberId,
			CreatedAt:  time.Now(),
		}
		var chatRoomMemberOppsite = model.ChatRoomMember{
			ChatRoomId:  chatRoom.ID,
			MemberId:    uint(toMemberId),
			IsNewToRead: 1,
			CreatedAt:   time.Now(),
		}
		if err := db.DB.Select("ChatRoomId", "MemberId", "CreatedAt").Create(&chatRoomMemberMe).Error; err != nil {
			return err
		}
		if err := db.DB.Select("ChatRoomId", "MemberId", "IsNewToRead", "CreatedAt").Create(&chatRoomMemberOppsite).Error; err != nil {
			return err
		}
		chatRoomId = chatRoom.ID
	} else {
		chatRoomId = existsChatRoom.ID
	}
	var chatRoomMemssage = model.ChatRoomMessage{
		ChatRoomId:   chatRoomId,
		FromMemberId: sendMessage.MemberId,
		ToMemberId:   uint(toMemberId),
		Message:      sendMessage.Message,
		CreatedAt:    time.Now(),
	}
	if err := db.DB.Select("ChatRoomId", "FromMemberId", "ToMemberId", "Message", "CreatedAt").Create(&chatRoomMemssage).Error; err != nil {
		return err
	}
	return nil
}

func GetUnreadCount(memberId uint) (int64, error) {
	return repository.GetUnreadCount(memberId)
}

func ReadAll(memberId uint) error {
	return repository.ReadAll(memberId)
}

// func SaveMember(member model.Member) error {
// 	return repository.SaveMember(member)
// }

// func GetThumbsUpList(pageNum, pageSize int, memberId uint, keyWord string) (*[]map[string]any, int64, error) {
// 	return repository.GetThumbsUpList(pageNum, pageSize, memberId, keyWord)
// }

// func UpdateThumbsUp(memberId, toMemberId uint, updateThumbsUp int) error {
// 	return repository.UpdateThumbsUp(memberId, toMemberId, updateThumbsUp)
// }

// func UpdateFavorite(memberId, toMemberId uint, updateFavorite int) error {
// 	return repository.UpdateFavorite(memberId, toMemberId, updateFavorite)
// }
