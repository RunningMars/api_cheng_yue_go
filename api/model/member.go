package model

import "time"

type Member struct {
	ID                      uint             `json:"id"`
	UserId                  uint             `json:"userId"`
	Mobile                  string           `json:"mobile"`
	Email                   string           `json:"email"`
	Password                string           `json:"password"`
	Sex                     int              `json:"sex"`
	Age                     int              `json:"age"`
	RealName                string           `json:"realName"`
	NickName                string           `json:"nickName"`
	Province                string           `json:"province"`
	City                    string           `json:"city"`
	Area                    string           `json:"area"`
	Height                  int              `json:"height"`
	Weight                  int              `json:"weight"`
	BodySize                string           `json:"bodySize"`
	School                  string           `json:"school"`
	WechatNo                string           `json:"wechatNo"`
	WechatMobile            string           `json:"wechatMobile"`
	EducationBackground     string           `json:"educationBackground"`
	EducationBackgroundCode int              `json:"educationBackgroundCode"`
	BirthYear               int              `json:"birthYear"`
	BirthMonth              int              `json:"birthMonth"`
	BirthDay                string           `json:"birthDay"`
	AnnualIncome            string           `json:"annualIncome"`
	MonthlyIncome           string           `json:"monthlyIncome"`
	AssetMoney              int              `json:"assetMoney"`
	AssetHouse              string           `json:"assetHouse"`
	AssetCar                string           `json:"assetCar"`
	MaritalStatus           string           `json:"maritalStatus"`
	ChildStatus             string           `json:"childStatus"`
	Vocation                string           `json:"vocation"`
	Job                     string           `json:"job"`
	AboutMe                 string           `json:"aboutMe"`
	Interest                string           `json:"interest"`
	HopeYou                 string           `json:"hopeYou"`
	IdentificationNo        string           `json:"identificationNo"`
	IdentificationName      string           `json:"identificationName"`
	IdentificationValidDate string           `json:"identificationValidDate"`
	ProfilePhoto            string           `json:"profilePhoto"`
	UploadProfilePhoto      string           `json:"uploadProfilePhoto"`
	AboutFamily             string           `json:"aboutFamily"`
	BrotherSister           string           `json:"brotherSister"`
	WantChild               string           `json:"wantChild"`
	WantMarry               string           `json:"wantMarry"`
	Ethnic                  string           `json:"ethnic"`
	Constellation           string           `json:"constellation"`
	SelfIntroduction        string           `json:"selfIntroduction"`
	MatingRequirement       string           `json:"matingRequirement"`
	AboutSmoke              string           `json:"aboutSmoke"`
	AboutDrink              string           `json:"aboutDrink"`
	IsAudit                 int              `json:"isAudit"`
	CreatedAt               time.Time        `json:"createdAt"`
	UpdatedAt               time.Time        `json:"updatedAt"`
	MemberImages            []MemberImage    `gorm:"foreignKey:MemberId;references:ID" json:"memberImages"`
	MemberRequest           MemberRequest    `gorm:"foreignKey:MemberId;references:ID" json:"memberRequest"`
	MemberThumbsUpToMember  []MemberThumbsUp `gorm:"foreignKey:ToMemberId;references:ID" json:"memberThumbsUpToMember"`
	MemberFavoriteToMember  []MemberFavorite `gorm:"foreignKey:ToMemberId;references:ID" json:"memberFavoriteToMember"`
	//DeletedAt     string        `json:"deleted_at"`
}

func SaveMember(member Member) {
	// 假设这是保存数据到数据库的代码
}

// 将 Member 的表名设置为 `member`
func (Member) TableName() string {
	return "member"
}
