package repository

import (
	"github.com/ZongweiBai/golang-in-action/config"
	// "github.com/jinzhu/gorm"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     int64
	Name   string
	Gender string
	Hobby  string
}

func SaveUserInfo() error {
	db := config.DBCONN
	uuid := config.IDNODE

	// 自动迁移
	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{uuid.Generate().Int64(), "七米", "男", "篮球"}
	u2 := UserInfo{uuid.Generate().Int64(), "沙河娜扎", "女", "足球"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	// 查询
	var u = new(UserInfo)
	db.First(u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)
	return nil
}
