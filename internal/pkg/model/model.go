package model

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"recruiting_system/internal/dep/db"
	"recruiting_system/internal/dep/log"
	"time"
)

type StudentInfo struct {
	gorm.Model
	StaffId        string `gorm:"index"`
	Name           string
	Phone          string
	Mail           string
	CommunityInfos []*CommunityInfo `gorm:"many2many:student_apply;"`
	Communities []*Community `gorm:"many2many:community_Principal;"`
}

//type communityPrincipal struct {
//	StaffId string
//	CommunityId string
//}

// Community 静态表

type Community struct {
	CommunityId string `gorm:"index" json:"communityId"`
	CommunityName string `json:"communityName"`
	CommunityInfo CommunityInfo `gorm:"foreignKey:CommunityId"`
	StudentInfos []*StudentInfo `gorm:"many2many:community_principal;"`
}

type CommunityInfo struct {
	//CommunityId     uint `gorm:"primarykey"`
	CommunityId     string `gorm:"index" json:"communityId"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	Name            string
	Introduction    string
	IsPublish       bool
	StudentInfos    StudentInfo `gorm:"many2many:student_apply;"`
	CommunityImages []CommunityImage `gorm:"foreignKey:ImageId"`
	CommunityJobs []CommunityJob `gorm:"foreignKey:JobId"`
}

type CommunityJob struct {
	JobId     uint `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	Name string
	Introduction string
}

type CommunityImage struct {
	ImageId   uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Url       string
}

func init() {
	db, err := db.ProvideDb("default")
	if err != nil {
		log.Fatal("pkg.db", zap.Error(err))
	}
	err = db.AutoMigrate(
		&StudentInfo{},
		&CommunityInfo{},
		&CommunityImage{},
	)
	if err != nil {
		log.Fatal("pkg.db", zap.Error(err))
	}
}
