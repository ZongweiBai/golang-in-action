package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	LOG    *zap.SugaredLogger
	CONFIG Config
	VIPER  *viper.Viper
	RDB    *redis.Client
)
