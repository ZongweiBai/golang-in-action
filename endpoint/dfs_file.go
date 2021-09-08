package endpoint

import (
	"fmt"
	CF "github.com/ZongweiBai/golang-in-action/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for index, file := range files {
		CF.LOG.Infof("开始上传文件：%s", file.Filename)
		dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d files uploaded!", len(files)),
	})
}
