package model

import "time"

type ChatRoomMember struct {
	ID          uint           `json:"id"`
	ChatRoomId  uint           `json:"chatRoomId"`
	MemberId    uint           `json:"memberId"`
	IsNewToRead uint8          `json:"isNewToRead"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	Member      map[string]any `gorm:"-" json:"member"`
}

// 配置表名
func (ChatRoomMember) TableName() string {
	return "chat_room_member"
}
