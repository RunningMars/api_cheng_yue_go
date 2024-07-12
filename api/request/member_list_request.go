package request

type MemberListReqeust struct {
	Id                             uint     `form:"id"`
	KeyWord                        string   `form:"keyWord"`
	Mobile                         string   `form:"mobile"`
	NickName                       string   `form:"nickName"`
	Sex                            int      `form:"-"` // 使用form - 标签忽略绑定
	MyMemberId                     uint     `form:"-"` // 使用form - 标签忽略绑定
	AgeMinRequest                  int      `form:"ageMinRequest"`
	AgeMaxRequest                  int      `form:"ageMaxRequest"`
	HeightMinRequest               int      `form:"heightMinRequest"`
	HeightMaxRequest               int      `form:"heightMaxRequest"`
	EducationBackgroundCodeRequest int      `form:"educationBackgroundCodeRequest"`
	AnnualIncomeRequest            string   `form:"annualIncomeRequest"`
	AnnualIncomeMinRequest         int      `form:"annualIncomeMinRequest"`
	AssetCarRequest                string   `form:"assetCarRequest"`
	AssetHouseRequest              []string `form:"assetHouseRequest[]"`
	MaritalStatusRequest           []string `form:"maritalStatusRequest[]"`
	WantChildRequest               string   `form:"wantChildRequest"`
	IsFavorite                     int      `form:"isFavorite"`
	IsThumbsUp                     int      `form:"isThumbsUp"`
}
