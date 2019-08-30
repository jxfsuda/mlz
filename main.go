package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sipt/GoJsoner"
	"io"
	"io/ioutil"
	"mlz/code"
	"mlz/conf"
	_ "mlz/docs" //初始化swagger
	"os"
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

	//初始化配置文件
	configFile:="conf/conf.json"
	jsonFile, err := ioutil.ReadFile(configFile)
	if err!=nil {
		panic("配置文件未找到,请查看"+configFile)
	}
	jsonStr ,err := GoJsoner.Discard(string(jsonFile))
	if err!=nil {
		panic("配置文件解析错误: "+err.Error())
	}

	conf.AppConfigObject = conf.AppConfig{}

	err =json.Unmarshal([]byte(jsonStr),&conf.AppConfigObject)
	if err!=nil {
		panic("配置文件解析错误: "+err.Error())
	}


	// 创建记录日志的文件
	f, _ := os.Create(conf.AppConfigObject.LogFile)
	//gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)





	//初始化路由设置,并启动服务
	code.InitRouters()














}
