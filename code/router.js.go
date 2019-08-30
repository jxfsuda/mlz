package code

import (
	"github.com/gin-gonic/gin"
	"mlz/code/controllers"
)

func InitRouters(router  *gin.Engine) {


// 接口组 无需授权的接口
	g:=router.Group("/api")
	{

		g.POST("/index", func(context *gin.Context) {
			c:= controllers.DemoController{}
			c.Index(context)
		})


	}



//接口组1  可以添加授权校验的 版本1的接口 , 支持内部继续分组
	g1:=router.Group("/api/v1")
	{

		demo:=g1.Group("/demo")
		{
			c:= controllers.DemoController{}
			demo.POST("/index", func(context *gin.Context) {
				c.Index(context)
			})
		}




	}







}
