package utils

import (
	"DouBanUpdater/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strconv"
	"time"
)

var GormDb *gorm.DB

func SQLInit() {
	var err error
	SQLServer := config.Config.SQLServer

	sqlStr := SQLServer.Username + ":" +
		SQLServer.Password + "@tcp(" +
		SQLServer.IP + ":" +
		strconv.Itoa(SQLServer.Port) + ")/" +
		SQLServer.Db + "?charset=" +
		SQLServer.Charset + "&parseTime=True&loc=" +
		SQLServer.Loc
	GormDb, err = gorm.Open(mysql.Open(sqlStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		//DisableForeignKeyConstraintWhenMigrating: true, // 禁止创建外键
		NamingStrategy: schema.NamingStrategy{ // 给创建表时候使用的
			SingularTable: true,
			// 全部的表名前面加前缀
			//TablePrefix: "mall_",
		},
	})
	sqlDB, err := GormDb.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		fmt.Println("数据库连接错误", err)
		return
	}
}
