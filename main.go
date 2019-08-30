package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sipt/GoJsoner"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io/ioutil"
	"mlz/code"
	"mlz/conf"
	_ "mlz/docs" //初始化swagger
	"strconv"
)

// @title Gin API123
// @version 1.0b125
// @description Golang API 演示, 主要是促成代码生成和数据库以及缓存使用,编写简单易于开发业务的框架 <br>技术栈:GOLang 1.12 ,  GIN , XORM , mysql , redis
// @termsOfService http://git.vs9.cn

// @contact.name jif
// @contact.url http://www.swagger.io/support
// @contact.email 35802713@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @ _host localhost
// @BasePath /
func main() {
	configFile:="conf/conf.json"
	jsonFile, err := ioutil.ReadFile(configFile)
	if err!=nil {
		panic("配置文件未找到,请查看"+configFile)
	}
	jsonStr ,err := GoJsoner.Discard(string(jsonFile))
	if err!=nil {
		panic("配置文件解析错误: "+err.Error())
	}

	appConfig:= conf.AppConfig{}

	err =json.Unmarshal([]byte(jsonStr),&appConfig)
	if err!=nil {
		panic("配置文件解析错误: "+err.Error())
	}




	r := gin.Default()

	if appConfig.RunMode=="dev"{
		//home,err:=conf.Home()
		//  //这样调用,只能下次进入的时候,才能看到最新文档,就当是再运行一次咯,暂时这样
		//cmd:= exec.Command(home+"/go/bin/swag","init")
		//err=cmd.Run()
		//if err!=nil {
		//	panic("配置文件解析错误: "+err.Error())
		//}else{
		//	log.Printf("执行命令swag init 成功")
		//}

		//注册swagger访问地址   /docs/index.html
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}else{
		gin.SetMode(gin.ReleaseMode)
	}


	//随便注册一个主页的路由, 因为是api项目,不展示网页. 前端用VUE即可
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is  API HOST",
		})
	})

	//初始化路由设置
	code.InitRouters(r)




	r.Run(":"+ strconv.Itoa(appConfig.WebConfig.Port)) // listen and serve on 0.0.0.0:8080
}
