package community

import (
	"fmt"
	"gorm.io/gorm"
	db1 "recruiting_system/internal/dep/db"
	"recruiting_system/internal/pkg/model"
)

var (
	db *gorm.DB
)

func init() {
	tempDb, err := db1.ProvideDb("default")
	if err != nil {
		panic(err)
	}
	db = tempDb
}

func GetAllCommunities() []*model.Community {
	communities := make([]*model.Community, 0, 100)
	// todo
	_ = db.Model(&model.Community{}).Find(&communities).Error
	return communities
}

func AddCommunityRegister(staffId string, communityId string) error {
	tempCommunity := new(model.Community)
	if err := db.
		Model(&model.Community{}).
		Where(&model.Community{CommunityId: communityId}).
		Find(tempCommunity).Error; err != nil {
		return fmt.Errorf("addCommunityRegister:%w", err)
	}

	if err := db.
		Model(&model.StudentInfo{}).
		Association("Communities").
		Append(&model.Community{CommunityId: communityId}); err != nil {
		return fmt.Errorf("addCommunityRegister:%w", err)
	}

	return nil
}

func GetCommunityIdByStaffId(staffId string) (communityId string, err error) {
	community := new(model.Community)
	err = db.Model(&model.StudentInfo{}).Association("Communities").Find(community)
	if err != nil {
		return "", err
	}
	return community.CommunityId, nil
}
