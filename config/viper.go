package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"github.com/ZongweiBai/learning-go/core"
)

// 初始化viper
func InitViper() *viper.Viper {
	viper := viper.New()
	viper.SetConfigFile("./config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(&core.CONFIG); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s", err))
	}

	// 监听文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生了变动", in.Name)
		if err := viper.Unmarshal(&core.CONFIG); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s", err))
		}
	})

	return viper
}