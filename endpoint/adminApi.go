package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/ZongweiBai/learning-go/repository"
	"github.com/ZongweiBai/learning-go/config"
)

func AdminHandler(c *gin.Context) {
	config.LOG.Debugf("进入到AdminHandler方法：%s", "李三四")
	c.JSON(200, &repository.User{ID: 20001, Name: "李三四"})
}