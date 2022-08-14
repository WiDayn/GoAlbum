package controller

import (
	"DouBanUpdater/model"
	"DouBanUpdater/response"
	"DouBanUpdater/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func GetList(c *gin.Context) {
	doubanId := c.Query("doubanId")
	startIndexText := c.DefaultQuery("startIndex", "0")

	var user model.User

	startIndex, _ := strconv.ParseInt(startIndexText, 10, 32)

	utils.GormDb.Model(&model.User{DoubanId: doubanId}).Preload("SubjectIndexes", func(db *gorm.DB) *gorm.DB {
		return db.Limit(16).Offset(int(startIndex)).Order("`index`")
	}).First(&user)

	var subjectIds []string
	var subjectIdString string

	for _, subjectIndex := range user.SubjectIndexes {
		subjectIds = append(subjectIds, subjectIndex.SubjectId)
		subjectIdString += "'" + subjectIndex.SubjectId + "' ,"
	}

	subjectIdString = subjectIdString[0 : len(subjectIdString)-2]

	var subjects []model.Subject

	utils.GormDb.Order("field(subject_id, "+subjectIdString+")").Find(&subjects, "subject_id IN ?", subjectIds)

	response.Success(c, gin.H{
		"startIndex": startIndex,
		"list":       subjects,
	}, "Success")
}
