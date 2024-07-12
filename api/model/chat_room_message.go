package model

import "time"

type ChatRoomMessage struct {
	ID           uint           `json:"id"`
	ChatRoomId   uint           `json:"chatRoomId"`
	FromMemberId uint           `json:"fromMemberId"`
	ToMemberId   uint           `json:"toMemberId"`
	Message      string         `json:"message"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	FromMember   map[string]any `gorm:"-" json:"fromMember"`
	ToMember     map[string]any `gorm:"-" json:"toMember"`
}

// 配置表名
func (ChatRoomMessage) TableName() string {
	return "chat_room_message"
}
