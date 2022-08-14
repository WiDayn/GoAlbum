package collector

import (
	"DouBanUpdater/logger"
	"strconv"
	"time"
)

func Run() {
	go defaultHandler()
}

func defaultHandler() {
	for true {
		MQSize := DefaultMessageQueue.Size()
		if MQSize != 0 {
			logger.Info("Now default MQ size："+strconv.Itoa(MQSize), nil)
			topMessage := DefaultMessageQueue.Top()
			logger.Info("Starting handle top:"+topMessage.SourceType+" "+topMessage.Id, nil)
			if topMessage.SourceType == "douban" {
				CollectDoubanUser(topMessage.Id)
			}
			logger.Info("Success handle! Now pop the top", nil)
			DefaultMessageQueue.Pop()
		}
		time.Sleep(time.Second * 1)
	}
}
