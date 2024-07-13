package repository

import (
	"api_go/api/model"
	"api_go/api/request"
	"api_go/api/util"
	"api_go/pkg/db"
	"fmt"
	"time"
	// "fmt"
)

func GetMembers(page, pageSize int, memberListReqeust request.MemberListReqeust) (*[]model.Member, int64, error) {
	var members []model.Member
	offset := (page - 1) * pageSize
	query := db.DB.Model(&model.Member{}).Preload("MemberImages").Offset(offset).Limit(pageSize)

	if memberListReqeust.KeyWord != "" {
		query = query.Where("nick_name LIKE ?", "%"+memberListReqeust.KeyWord+"%")
	}
	if memberListReqeust.Sex > 0 {
		query = query.Where("sex = ?", memberListReqeust.Sex)
	}
	if memberListReqeust.MyMemberId > 0 {
		query = query.Where("id <> ?", memberListReqeust.MyMemberId)
	}
	if memberListReqeust.AgeMinRequest > 0 {
		query = query.Where("age >= ?", memberListReqeust.AgeMinRequest)
	}
	if memberListReqeust.AgeMaxRequest > 0 {
		query = query.Where("age <= ?", memberListReqeust.AgeMaxRequest)
	}
	if memberListReqeust.HeightMinRequest > 0 {
		query = query.Where("height >= ?", memberListReqeust.HeightMinRequest)
	}
	if memberListReqeust.HeightMaxRequest > 0 {
		query = query.Where("height <= ?", memberListReqeust.HeightMaxRequest)
	}
	if memberListReqeust.EducationBackgroundCodeRequest > 0 {
		query = query.Where("education_background_code >= ?", memberListReqeust.EducationBackgroundCodeRequest)
	}
	if memberListReqeust.AnnualIncomeRequest != "" {
		query = query.Where("annual_income = ?", memberListReqeust.AnnualIncomeRequest)
	}
	if memberListReqeust.AnnualIncomeMinRequest > 0 {
		query = query.Where("annual_income_min >= ?", memberListReqeust.AnnualIncomeMinRequest)
	}
	if memberListReqeust.AssetCarRequest != "" {
		query = query.Where("asset_car = ?", memberListReqeust.AssetCarRequest)
	}
	if len(memberListReqeust.AssetHouseRequest) > 0 {
		query = query.Where("asset_house IN ?", memberListReqeust.AssetHouseRequest)
	}
	if len(memberListReqeust.MaritalStatusRequest) > 0 {
		query = query.Where("marital_status IN ?", memberListReqeust.MaritalStatusRequest)
	}
	if memberListReqeust.WantChildRequest != "" {
		query = query.Where("child_status = ?", memberListReqeust.WantChildRequest)
	}
	if memberListReqeust.IsFavorite == 1 {
		subQuery := db.DB.Model(&model.MemberFavorite{}).Select("id").Where("member_favorite.to_member_id = member.id").
			Where("member_favorite.member_id = ?", memberListReqeust.MyMemberId)
		query = query.Where("EXISTS (?)", subQuery)
	}
	if memberListReqeust.IsThumbsUp == 1 {
		subQuery := db.DB.Model(&model.MemberThumbsUp{}).Select("id").Where("member_thumbs_up.to_member_id = member.id").
			Where("member_thumbs_up.member_id = ?", memberListReqeust.MyMemberId)
		query = query.Where("EXISTS (?)", subQuery)
	}
	var total int64

	if err := query.Count(&total).Find(&members).Error; err != nil {
		return nil, 0, err
	}
	return &members, total, nil
}

func GetMemberByUserId(userId uint) (*model.Member, error) {
	var member model.Member
	if err := db.DB.Where("user_id = ?", userId).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func GetMemberById(memberId uint, currentMemberId uint) (*model.Member, error) {
	var member model.Member
	if err := db.DB.Debug().Model(&model.Member{}).
		Preload("MemberImages").
		Preload("MemberRequest").
		Preload("MemberThumbsUpToMember", "member_thumbs_up.member_id = ?", currentMemberId).
		Preload("MemberFavoriteToMember", "member_favorite.member_id = ?", currentMemberId).
		Where("member.id = ?", memberId).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func GetMemberById2(memberId uint) (*model.Member, error) {
	var member model.Member
	if err := db.DB.Debug().Model(&model.Member{}).
		Where("member.id = ?", memberId).Find(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}
func GetMemberByIds(memberIds []uint) (*[]model.Member, error) {
	var members []model.Member
	if err := db.DB.Debug().Model(&model.Member{}).
		Where("member.id IN ?", memberIds).
		Find(&members).Error; err != nil {
		return nil, err
	}
	return &members, nil
}

func SaveMember(member model.Member) error {

	// DELETE from member_image where member_id = ?;
	db.DB.Where("member_id = ?", member.ID).Delete(&model.MemberImage{})

	if len(member.MemberImages) > 0 {
		// save member images
		result := db.DB.Create(member.MemberImages)
		if result.Error != nil {
			return result.Error
		}
	}

	// 解析出生日期
	birthdate, err := time.Parse("2006-01-02", member.BirthDay)
	if err == nil {
		// 计算年龄
		age := util.CalculateAge(birthdate)
		if age > 0 {
			member.Age = age
		} else {
			member.Age = 0
		}
	}

	// 根据 `struct` 更新属性，只会更新非零值的字段
	db.DB.Model(&member).Updates(member)

	db.DB.Model(&member.MemberRequest).Updates(member.MemberRequest)

	return nil
}

func GetThumbsUpList(pageNum, pageSize int, memberId uint, keyWord string) (*[]map[string]any, int64, error) {
	var thumbsUpList []model.MemberThumbsUp
	offset := (pageNum - 1) * pageSize
	query := db.DB.Model(&model.MemberThumbsUp{}).Offset(offset).Limit(pageSize)
	query = query.Where("to_member_id = ?", memberId)

	var total int64

	if err := query.Count(&total).Find(&thumbsUpList).Error; err != nil {
		return nil, 0, err
	}

	// 提取所有的 member_id
	memberIds := make([]uint, len(thumbsUpList))
	for i, thumbsUp := range thumbsUpList {
		memberIds[i] = thumbsUp.MemberId
	}

	// 批量查询 member 信息
	var members []model.Member
	if err := db.DB.Where("id IN ?", memberIds).
		Select("id,user_id,nick_name,profile_photo,age").Find(&members).Error; err != nil {
		return nil, 0, err
	}

	// 将 member 信息写入到 thumbsUpList 中
	memberMap := make(map[uint]model.Member)
	for _, member := range members {
		memberMap[member.ID] = member
	}
	var res []map[string]any
	res = make([]map[string]any, len(thumbsUpList))
	for i, thumbsUp := range thumbsUpList {
		if member, found := memberMap[thumbsUp.MemberId]; found {
			res[i] = map[string]any{
				"id":          thumbsUp.ID,
				"memberId":    thumbsUp.MemberId,
				"toMemeberId": thumbsUp.ToMemberId,
				"isThumbsUp":  thumbsUp.IsThumbsUp,
				"createdAt":   thumbsUp.CreatedAt,
				"updatedAt":   thumbsUp.UpdatedAt,
				"member": map[string]any{
					"id":           member.ID,
					"userId":       member.UserId,
					"nickName":     member.NickName,
					"profilePhoto": member.ProfilePhoto,
					"age":          member.Age,
				},
			}
		}
	}
	fmt.Println("res", res)
	return &res, total, nil
}

func UpdateThumbsUp(memberId, toMemberId uint, updateThumbsUp int) error {
	if updateThumbsUp == 0 {

		db.DB.Where("member_id = ?", memberId).Where("to_member_id = ?", toMemberId).Delete(&model.MemberThumbsUp{})
	} else {
		var existMemberThumbsUp model.MemberThumbsUp
		db.DB.Model(&model.MemberThumbsUp{}).
			Where("member_id = ?", memberId).
			Where("to_member_id = ?", toMemberId).Limit(1).
			Find(&existMemberThumbsUp)

		if existMemberThumbsUp.ID == 0 {
			memberThumbsUp := model.MemberThumbsUp{
				MemberId:   memberId,
				ToMemberId: toMemberId,
				CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
			}
			if err := db.DB.Select("MemberId", "ToMemberId", "CreatedAt").Create(&memberThumbsUp).Error; err != nil {
				// 处理创建错误
				fmt.Println("创建错误: ", err)
				return err
			}
		}
	}
	return nil
}

func UpdateFavorite(memberId, toMemberId uint, updateFavorite int) error {
	if updateFavorite == 0 {
		db.DB.Where("member_id = ?", memberId).Where("to_member_id = ?", toMemberId).Delete(&model.MemberFavorite{})
	} else {
		var existMemberFavorite model.MemberFavorite
		db.DB.Model(&model.MemberFavorite{}).
			Where("member_id = ?", memberId).
			Where("to_member_id = ?", toMemberId).Limit(1).
			Find(&existMemberFavorite)
		if existMemberFavorite.ID == 0 {
			memberFavorite := model.MemberFavorite{
				MemberId:   memberId,
				ToMemberId: toMemberId,
				CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
			}
			if err := db.DB.Select("MemberId", "ToMemberId", "CreatedAt").Create(&memberFavorite).Error; err != nil {
				// 处理创建错误
				fmt.Println("创建错误: ", err)
				return err
			}
		}
	}
	return nil
}
