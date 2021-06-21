package main

import (
	// "fmt"
	"github.com/ZongweiBai/learning-go/config"
	"github.com/ZongweiBai/learning-go/server"
	"github.com/ZongweiBai/learning-go/core"
	_ "github.com/ZongweiBai/learning-go/docs"
	"github.com/ZongweiBai/learning-go/repository"
	"go.uber.org/zap"
	"log"
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

	core.VIPER = config.InitViper()

	var zapLogger *zap.Logger
	zapLogger, core.LOG = config.InitLogger()

	server.InitWebServer(zapLogger)

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)

	userOne := repository.NewUser(1, "李四")
	log.Printf("用户ID: %v, 用户名：%s", userOne.ID, userOne.Name)
	userOne.ChangeId(12, userOne)
	userOne.ChangeName("李四五")
	log.Printf("用户ID: %v, 用户名：%s", userOne.ID, userOne.Name)

}