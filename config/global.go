package config

import (
	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	LOG    *zap.SugaredLogger
	IDNODE *snowflake.Node
	CONFIG Config
	VIPER  *viper.Viper
	RDB    *redis.Client
	DBCONN *gorm.DB
)

// InitSnowflake 初始化snowflake
func InitSnowflake() {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		LOG.Error("初始化snowflake失败", err)
		return
	}
	IDNODE = node
}
