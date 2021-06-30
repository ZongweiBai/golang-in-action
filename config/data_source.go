package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"sync"
)

var dsOnce sync.Once

func InitDbConn() {
	dsOnce.Do(func() {
		var dialet = "postgres"
		if (CONFIG.DataSource.Dialect != "") {
			dialet = CONFIG.DataSource.Dialect
		}
		db, err := gorm.Open(dialet, CONFIG.DataSource.Url)
		if err != nil {
			LOG.Errorf("初始化DB Connection失败", err)
			panic(err)
		}

		//最大连接数
		db.DB().SetMaxOpenConns(CONFIG.DataSource.MaxPoolSize)
		//最大空闲连接数
		db.DB().SetMaxIdleConns(CONFIG.DataSource.MinIdle)
		db.LogMode(true)

		// set schema here.
		gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
			if CONFIG.DataSource.Scheme != "" {
				return CONFIG.DataSource.Scheme + "." + tableName
			} else {
				return tableName
			}
		}

		DBCONN = db
	})
}
