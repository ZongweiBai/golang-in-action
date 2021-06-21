package server

import (
	_ "github.com/ZongweiBai/learning-go/docs"
	"github.com/ZongweiBai/learning-go/endpoint"
	"github.com/ZongweiBai/learning-go/repository"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
	"log"
	"strconv"
	"time"
)

func InitWebServer(zapLogger *zap.Logger) {
	// r := gin.Default()
	r := gin.New()
	r.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zapLogger, true))

	r.Use(costTime())

	// http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		json := repository.User{}
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

	// 获取Token
	r.POST("/oauth/token", endpoint.GenerateAcessToken)
	// 校验Token
	r.GET("/oauth/token/validate", endpoint.ValidateAcessToken)

	users := []repository.User{{ID: 1, Name: "张三"}, {ID: 2, Name: "李四"}}
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		idUInt, _ := strconv.ParseUint(id, 0, 64)
		c.JSON(200, &repository.User{ID: idUInt, Name: "张三" + id})
	})

	r.GET("/users/query", func(c *gin.Context) {
		id := c.DefaultQuery("id", "0")
		idUInt, _ := strconv.ParseUint(id, 0, 64)
		c.JSON(200, &repository.User{ID: idUInt, Name: c.Query("name")})
	})

	// 路由嵌套分组
	v1Group := r.Group("/v1")
	v1Group.Use(gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))
	{

		v1Group.GET("/users", func(c *gin.Context) {
			c.JSON(200, &repository.User{ID: 10001, Name: "李四"})
		})

		v1AdminGroup := v1Group.Group("/admin")
		{
			// v1AdminGroup.GET("/users", func(c *gin.Context) {
			// 	c.JSON(200, &User{ID: 20001, Name: "李四"})
			// })
			v1AdminGroup.GET("/users", endpoint.AdminHandler)
		}
	}
	r.Run(":8080")
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