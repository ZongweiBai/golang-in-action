package main

import (
	// "fmt"
	"github.com/ZongweiBai/learning-go/config"
	"github.com/ZongweiBai/learning-go/core"
	"github.com/ZongweiBai/learning-go/task"
	_ "github.com/ZongweiBai/learning-go/docs"
	"github.com/ZongweiBai/learning-go/repository"
	"go.uber.org/zap"
)

// @title Learning-Go Swagger文档
// @version 1.0
// @description Go入门学习项目
// @termsOfService https://github.com/ZongweiBai

// @contact.name ZongweiBai
// @contact.url https://github.com/ZongweiBai
// @contact.email zongwei.bai@gmail.com

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @host 127.0.0.1:8080
// @BasePath /
func main() {

	config.VIPER = config.InitViper()

	var zapLogger *zap.Logger
	zapLogger, config.LOG = config.InitLogger()

	// 初始化定时任务
	task.SetupTasks()

	// 初始化tcp socket服务
	go core.InitSocketServer()

	// 初始化web服务器
	core.InitWebServer(zapLogger)

	userOne := repository.NewUser(1, "李四")
	config.LOG.Debugf("用户ID: %v, 用户名：%s", userOne.ID, userOne.Name)
	userOne.ChangeId(12, userOne)
	userOne.ChangeName("李四五")
	config.LOG.Debugf("用户ID: %v, 用户名：%s", userOne.ID, userOne.Name)

}
