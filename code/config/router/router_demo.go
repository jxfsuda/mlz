package router

import (
	"mlz/code/controllers/Demo"
)

func route_demo (){


	//接口组1  可以添加授权校验的 版本1的接口 , 支持内部继续分组
	g1:= router.Group("/api/demo")
	{

		demo:=g1.Group("/v1")
		{
			demo.POST("/index", Demo.Index)
		}




	}

}