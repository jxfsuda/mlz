package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sipt/GoJsoner"
	"io"
	"io/ioutil"
	"log"
	"mlz/conf"
	"mlz/conf/router"
	_ "mlz/docs" //初始化swagger
	"mlz/iolib/xorm"
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

	eng,err:=xorm.NewMysqlEngine(conf.AppConfigObject.DataSource)
	if err!=nil {
		panic("数据库连接失败: "+err.Error())
	}else{
		log.Println("数据库初始化成功")
	}

	conf.AppConfigObject.Db = eng


	//初始化路由设置,并启动服务
	router.InitRouters()







/**
打包
  go get github.com/mitchellh/gox

gox   默认会打包所有平台
  gox -os="linux"
  gox -osarch="linux/amd64"

  gox -osarch="darwin/amd64"  -ldflags="-s -w"

   macos 上通过 brew 安装brew install upx
   upx -9 -o ./mlz ./mlz_darwin_amd64    -9 压缩级别  -o 压缩后的文件名   最后一个参数 要压缩的文件


    Mac 下打包 Linux 环境下运行的程序
  env GOOS=linux GOARCH=386 go build main.go

 */






}
