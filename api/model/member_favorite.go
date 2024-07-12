package model

type MemberFavorite struct {
	ID         uint   `json:"id"`
	MemberId   uint   `json:"memberId"`
	ToMemberId uint   `json:"toMemberId"`
	IsFavorite int    `json:"isFavorite"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

func (MemberFavorite) TableName() string {
	return "member_favorite"
}
