package controller

import (
	"DouBanUpdater/collector"
	"DouBanUpdater/model"
	"DouBanUpdater/response"
	"DouBanUpdater/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func UpdateByDoubanId(c *gin.Context) {
	doubanId := c.Query("doubanId")

	var user model.User

	utils.GormDb.Model(&model.User{DoubanId: doubanId}).First(&user)

	if user.UpdatedAt.After(time.Now().Add(-time.Second * 5)) {
		response.Fail(c, nil, "Requests are too frequent")
		return
	}

	if collector.DefaultMessageQueue.Contains(collector.DefaultRequestMessage{SourceType: "douban", Id: doubanId}) {
		response.Fail(c, nil, "Already handling")
		return
	}

	collector.DefaultMessageQueue.Add(collector.DefaultRequestMessage{SourceType: "douban", Id: doubanId})
	response.Success(c, nil, "Starting...")
}
