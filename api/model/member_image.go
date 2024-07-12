package model

type MemberImage struct {
	ID        uint    `json:"id"`
	MemberId  uint   `json:"memberId"`
	Url       string `json:"url"`
	UploadUrl string `json:"uploadUrl"`
	CreatedAt string `json:"createdAt"`
	CreatedBy int    `json:"createdBy"`
	UpdatedAt string `json:"updatedAt"`
	UpdatedBy int    `json:"updatedBy"`
}

// 将 Member 的表名设置为 `member`
func (MemberImage) TableName() string {
	return "member_image"
}
