package model

import "time"

type ChatRoom struct {
	ID                     uint             `json:"id"`
	ChatRoomName           string           `json:"chatRoomName"`
	CreatedAt              time.Time        `json:"createdAt"`
	UpdatedAt              time.Time        `json:"updatedAt"`
	ChatRoomOppositeMember []ChatRoomMember `gorm:"foreignKey:ChatRoomId;references:ID" json:"chatRoomOppositeMember"`
	ChatRoomMeMember       []ChatRoomMember `gorm:"foreignKey:ChatRoomId;references:ID" json:"chatRoomMeMember"`
	ChatRoomLastMessage    *ChatRoomMessage  `json:"chatRoomLastMessage"`
}

// 配置表名
func (ChatRoom) TableName() string {
	return "chat_room"
}
