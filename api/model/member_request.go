package model

import "time"

type MemberRequest struct {
	ID                         uint       `json:"id"`
	MemberId                   uint      `json:"memberId"`
	AgeMinRequest              int       `json:"ageMinRequest"`
	AgeMaxRequest              int       `json:"ageMaxRequest"`
	CityRequest                string    `json:"cityRequest"`
	HeightMinRequest           int       `json:"heightMinRequest"`
	HeightMaxRequest           int       `json:"heightMaxRequest"`
	WeightMinRequest           int       `json:"weightMinRequest"`
	WeightMaxRequest           int       `json:"weightMaxRequest"`
	BodySizeRequest            string    `json:"bodySizeRequest"`
	EducationBackgroundRequest string    `json:"educationBackgroundRequest"`
	AnnualIncomeRequest        int       `json:"annualIncomeRequest"`
	AssetHouseRequest          string    `json:"assetHouseRequest"`
	AssetCarRequest            string    `json:"assetCarRequest"`
	MaritalStatusRequest       string    `json:"maritalStatusRequest"`
	ChildStatusRequest         string    `json:"childStatusRequest"`
	JobRequest                 string    `json:"jobRequest"`
	AboutFamilyRequest         string    `json:"aboutFamilyRequest"`
	BrotherSisterRequest       string    `json:"brotherSisterRequest"`
	WantChildRequest           string    `json:"wantChildRequest"`
	WantMarryRequest           string    `json:"wantMarryRequest"`
	AboutSmokeRequest          string    `json:"aboutSmokeRequest"`
	AboutDrinkRequest          string    `json:"aboutDrinkRequest"`
	CreatedAt                  time.Time `json:"createdAt"`
	UpdatedAt                  time.Time `json:"updatedAt"`
}

// 将 Member 的表名设置为 `member`
func (MemberRequest) TableName() string {
	return "member_request"
}
