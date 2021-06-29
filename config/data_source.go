package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDbConn() {
	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	if err != nil {
		LOG.Errorf("初始化DB Connection失败", err)
		panic(err)
	}
	defer db.Close()

	//最大连接数
	db.DB().SetMaxOpenConns(8)
	//最大空闲连接数
	db.DB().SetMaxIdleConns(2)
	db.LogMode(true)

	DBCONN = db
}
