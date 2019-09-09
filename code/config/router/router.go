package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"mlz/code/config"
	_ "mlz/code/validator"
	"mlz/iolib/gin_ext"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

var router = gin.New()

func InitRouters() {


	// 全局中间件
	// 使用 Logger 中间件
	router.Use(gin.Logger())

	// 使用 Recovery 中间件
	router.Use(gin_ext.Recovery())

	router.Delims("{%", "%}")

	router.LoadHTMLGlob("conf/templates/*")



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










	if config.AppConfigObject.RunMode=="dev"{
		//注册swagger访问地址   /docs/index.html
		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// 调试模式 ,查看内存占用之类的
		if config.AppConfigObject.DebugPort==0 {
			config.AppConfigObject.DebugPort=7234}
		go func() {
			log.Fatal(http.ListenAndServe("0.0.0.0:"+ strconv.Itoa(config.AppConfigObject.DebugPort), nil))
		}()

	}else{
		gin.SetMode(gin.ReleaseMode)
	}


	//随便注册一个主页的路由, 因为是api项目,不展示网页. 前端用VUE即可
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"res":""})
	})

	////扩展路由配置 当前文件用于入口处理,以及通用底层配置,其他文件可以一个模块一个文件,避免团队开发冲突
	// 实际开发中,在这里添加一个方法就表示在同级目录增加一个文件和方法( 每个文件配置一个路径,名称为 router_{module_name} ) ,用于区分和查找
	route_index()
	route_demo()





	router.Run(":"+ strconv.Itoa(config.AppConfigObject.WebConfig.Port)) // listen and serve on 0.0.0.0:8080
}
