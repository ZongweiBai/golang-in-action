package core

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	LOG *zap.SugaredLogger
	CONFIG Config
	VIPER *viper.Viper
)