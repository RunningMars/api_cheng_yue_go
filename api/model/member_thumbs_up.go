package model

type MemberThumbsUp struct {
	ID         uint    `json:"id"`
	MemberId   uint   `json:"memberId"`
	ToMemberId uint   `json:"toMemberId"`
	IsThumbsUp int    `json:"isThumbsUp"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	//Member     Member `json:"member"`
}

func (MemberThumbsUp) TableName() string {
	return "member_thumbs_up"
}
