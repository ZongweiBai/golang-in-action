package core

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/ZongweiBai/golang-in-action/config"
	_ "github.com/ZongweiBai/golang-in-action/docs"
	"github.com/ZongweiBai/golang-in-action/endpoint"
	"github.com/ZongweiBai/golang-in-action/repository"
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
)

func InitWebServer(zapLogger *zap.Logger) {
	// r := gin.Default()
	r := gin.New()
	r.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zapLogger, true))
	// use prometheus metrics exporter middleware.
	//
	// ginprom.PromMiddleware() expects a ginprom.PromOpts{} poniter.
	// It is used for filtering labels by regex. `nil` will pass every requests.
	//
	// ginprom promethues-labels:
	//   `status`, `endpoint`, `method`
	//
	// for example:
	// 1). I don't want to record the 404 status request. That's easy for it.
	// ginprom.PromMiddleware(&ginprom.PromOpts{ExcludeRegexStatus: "404"})
	//
	// 2). And I wish to ignore endpoints started with `/prefix`.
	// ginprom.PromMiddleware(&ginprom.PromOpts{ExcludeRegexEndpoint: "^/prefix"})
	r.Use(ginprom.PromMiddleware(nil))
	r.Use(costTime())

	// 性能监控
	// http://localhost:8080/debug/pprof
	pprof.Register(r)

	// register the `/metrics` route.
	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))

	// swagger docs
	// http://localhost:8080/api-docs/swagger/index.html
	r.GET("/api-docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 路由嵌套分组
	v1Group := r.Group("/v1")
	{
		// oauth路由
		oauthGroup := v1Group.Group("/oauth")
		oauthGroup.Use(gin.BasicAuth(gin.Accounts{
			"admin": "123456",
		}))
		{
			// 获取Token
			oauthGroup.POST("/token", endpoint.GenerateAcessToken)
			// 校验Token
			oauthGroup.GET("/token/validate", endpoint.ValidateAcessToken)
		}

		v1Group.GET("/users", func(c *gin.Context) {
			c.JSON(200, &repository.User{ID: 10001, Name: "李四"})
		})

		v1AdminGroup := v1Group.Group("/admin")
		{
			v1AdminGroup.GET("/users", endpoint.AdminHandler)
		}
	}

	testGroup := r.Group("/test")
	{
		// http://localhost:8080/?media=123&media=345&ids[a]=123&ids[b]=456&ids[c]=789
		testGroup.GET("/", func(c *gin.Context) {
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
		testGroup.POST("/", func(c *gin.Context) {
			wechat := c.PostForm("wechat")
			c.JSON(200, gin.H{
				"wechat": wechat,
			})
		})

		// 获取JSON body  二选一
		testGroup.POST("/users", func(c *gin.Context) {
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

		users := []repository.User{{ID: 1, Name: "张三"}, {ID: 2, Name: "李四"}}
		testGroup.GET("/users", func(c *gin.Context) {
			c.JSON(200, users)
		})

		testGroup.GET("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			idUInt, _ := strconv.ParseUint(id, 0, 64)
			c.JSON(200, &repository.User{ID: idUInt, Name: "张三" + id})
		})

		testGroup.GET("/users/query", func(c *gin.Context) {
			id := c.DefaultQuery("id", "0")
			idUInt, _ := strconv.ParseUint(id, 0, 64)
			c.JSON(200, &repository.User{ID: idUInt, Name: c.Query("name")})
		})
	}

	// r.Run(":8080")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	// 优雅关机
	shutdownGraceful(srv)
}

func shutdownGraceful(srv *http.Server) {
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	config.LOG.Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		config.LOG.Fatal("Server Shutdown: ", err)
	}
	config.LOG.Info("Server exiting")
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
		config.LOG.Debugf("the request URL %s cost %v", url, costTime)
	}
}
