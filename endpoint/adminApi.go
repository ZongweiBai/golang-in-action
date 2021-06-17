package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/ZongweiBai/learning-go/repository"
)

func AdminHandler(c *gin.Context) {
	c.JSON(200, &repository.User{ID: 20001, Name: "李三四"})
}