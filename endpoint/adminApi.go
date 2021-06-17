package endpoint

import (
	"github.com/gin-gonic/gin"
	"main"
)

func adminHandler(c *gin.Context) {
	c.JSON(200, &main.User{ID: 20001, Name: "李四"})
}