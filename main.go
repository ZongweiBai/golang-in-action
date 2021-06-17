package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"strconv"
	"endpoint"
)

func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)

	r := gin.Default()

	r.Use(gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))

	r.Use(costTime())

	// http://localhost:8080/?media=123&media=345&ids[a]=123&ids[b]=456&ids[c]=789
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":   "www.flysnow.org",
			"wechat": "flysnow_org",
			// 获取array参数
			"param": c.QueryArray("media"),
			// 获取map参数
			"ids": c.QueryMap("ids"),
		})
	})

	// 获取表单参数
	r.POST("/", func(c *gin.Context) {
		wechat := c.PostForm("wechat")
		c.JSON(200, gin.H{
			"wechat": wechat,
		})
	})

	// 获取JSON body  二选一
	r.POST("/users", func(c *gin.Context) {
		// json1 := make(map[string]interface{}) //注意该结构接受的内容
		// c.BindJSON(&json1)
		// fmt.Printf("%v", &json1)

		// var json User
		// 或者
		json := User{}
		err := c.BindJSON(&json)
		if err != nil {
			log.Println("===================>>>>>>>>")
			log.Println(err)
			log.Println("===================>>>>>>>>")
			c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		} else {
			c.JSON(200, json)
		}
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

	// 路由嵌套分组
	v1Group := r.Group("/v1")
	{
		v1Group.GET("/users", func(c *gin.Context) {
			c.JSON(200, &User{ID: 10001, Name: "李四"})
		})

		v1AdminGroup := v1Group.Group("/admin")
		{
			v1AdminGroup.GET("/users", func(c *gin.Context) {
				c.JSON(200, &User{ID: 20001, Name: "李四"})
			})
		}
	}
	r.Run(":8080")
}

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}


func costTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求当前时间
		nowTime := time.Now()

		// 请求处理
		c.Next()

		// 处理后获取耗时
		costTime := time.Since(nowTime)
		url := c.Request.URL.String()
		log.Printf("the request URL %s cost %v", url, costTime)
	}
}