package model

import "time"

type User struct {
	ID        uint      `json:"id"`
	Mobile    string    `json:"mobile"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// DeletedAt gorm.DeletedAt `gorm:"-"` // 通过 struct tag 移除
}

func SaveUser(user User) {
	// 假设这是保存数据到数据库的代码
}

// 将 User 的表名设置为 `user`
func (User) TableName() string {
	return "user"
}
