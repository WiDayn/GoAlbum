package model

import (
	"DouBanUpdater/utils"
	"gorm.io/gorm"
	"time"
)

type User struct {
	DoubanId       string `gorm:"type:varchar(20); primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	SubjectIndexes []SubjectIndex `gorm:"foreignKey:DoubanId"`
}

type SubjectIndex struct {
	DoubanId  string `gorm:"type:varchar(20); primary_key"`
	Index     int    `gorm:"primary_key"`
	SubjectId string `gorm:"type:varchar(20);"`
}

type Subject struct {
	SubjectId    string `gorm:"type:varchar(20); primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	ReleaseDate  string
	SubjectName  string
	SubjectPhoto string
}

func UserInit() {
	if err := utils.GormDb.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	if err := utils.GormDb.AutoMigrate(&SubjectIndex{}); err != nil {
		panic(err)
	}
	if err := utils.GormDb.AutoMigrate(&Subject{}); err != nil {
		panic(err)
	}
}
