package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":   "www.flysnow.org",
			"wechat": "flysnow_org",
		})
	})

	users := []User{{ID: 1, Name: "张三"}, {ID: 2, Name: "李四"}}
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		idUInt, _ := strconv.ParseUint(id, 0, 64)
		c.JSON(200, &User{ID: idUInt, Name: "张三" + id})
	})

	r.GET("/users/query", func(c *gin.Context) {
		id := c.DefaultQuery("id", "0")
		idUInt, _ := strconv.ParseUint(id, 0, 64)
		c.JSON(200, &User{ID: idUInt, Name: c.Query("name")})
	})
	r.Run(":8080")
}

type User struct {
	ID   uint64
	Name string
}
