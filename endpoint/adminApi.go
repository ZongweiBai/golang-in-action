package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/ZongweiBai/learning-go/repository"
	"github.com/ZongweiBai/learning-go/core"
	_ "github.com/ZongweiBai/learning-go/model"
)

// GetAdminHandler 获取Admin信息
// @Summary 获取Admin信息
// @Description 通过名称获取Admin信息
// @Tags Admin相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Basic 用户令牌"
// @Security model.ApiBacisAuth
// @Success 200 {object} repository.User
// @Router /v1/admin/users [get]
func AdminHandler(c *gin.Context) {
	core.LOG.Debugf("进入到AdminHandler方法：%s", "李三四")
	c.JSON(200, &repository.User{ID: 20001, Name: "李三四"})
}