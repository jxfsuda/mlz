package code

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"mlz/code/controllers"
	_ "mlz/code/validator"
	"mlz/conf"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

func InitRouters() {
	router := gin.Default()


	// LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			//这里一个坑,golang的日期格式化,必须是这个串... 见 https://www.jianshu.com/p/c7f7fbb16932
			param.TimeStamp.Format("2006-01-02 15:04:05.000"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))








	c:= controllers.DemoController{}

	if conf.AppConfigObject.RunMode=="dev"{
		//注册swagger访问地址   /docs/index.html
		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// 调试模式 ,查看内存占用之类的
		if conf.AppConfigObject.DebugPort==0 {conf.AppConfigObject.DebugPort=7234}
		go func() {
			log.Fatal(http.ListenAndServe("0.0.0.0:"+ strconv.Itoa(conf.AppConfigObject.DebugPort), nil))
		}()

	}else{
		gin.SetMode(gin.ReleaseMode)
	}


	//随便注册一个主页的路由, 因为是api项目,不展示网页. 前端用VUE即可
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is  API HOST",
		})
	})

// 接口组 无需授权的接口
	g:=router.Group("/api")
	{

		g.POST("/index", func(context *gin.Context) {

			c.Index(context)
		})


	}



//接口组1  可以添加授权校验的 版本1的接口 , 支持内部继续分组
	g1:=router.Group("/api/v1")
	{

		demo:=g1.Group("/demo")
		{

			demo.POST("/index", func(context *gin.Context) {
				c.Index(context)
			})
		}




	}






	router.Run(":"+ strconv.Itoa(conf.AppConfigObject.WebConfig.Port)) // listen and serve on 0.0.0.0:8080
}
