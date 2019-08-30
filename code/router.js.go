package code

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"mlz/code/controllers"
	"mlz/conf"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

func InitRouters() {
	router := gin.Default()
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
