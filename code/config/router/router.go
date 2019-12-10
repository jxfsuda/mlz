package router

import (
	"context"
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
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var router = gin.New()

func InitRouters() {



	//"github.com/thinkerou/favicon"
///	router.Use(favicon.New("./favicon.ico"))
	// 全局中间件
	// 使用 Logger 中间件
	router.Use(gin.Logger())

	// 使用 Recovery 中间件
	router.Use(gin_ext.Recovery())
	router.Use(Cors())
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
		c.HTML(200, "index.html", gin.H{"title":"这个字符串是后端发来的"})
	})

	////扩展路由配置 当前文件用于入口处理,以及通用底层配置,其他文件可以一个模块一个文件,避免团队开发冲突
	// 实际开发中,在这里添加一个方法就表示在同级目录增加一个文件和方法( 每个文件配置一个路径,名称为 router_{module_name} ) ,用于区分和查找
	route_index()
	route_demo()
	route_su()
	route_scrapy ()

	server := &http.Server{
		Addr:           ":"+strconv.Itoa(config.AppConfigObject.WebConfig.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go server.ListenAndServe()

	// 设置优雅退出
	gracefulExitWeb(server)



}

func gracefulExitWeb(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	fmt.Println("got a signal", sig)
	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil{
		fmt.Println("err", err)
	}

	// 看看实际退出所耗费的时间
	fmt.Println("------exited--------", time.Since(now))
}


func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}

}