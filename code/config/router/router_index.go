package router

import (
	"mlz/code/controllers/Demo"
)

func  route_index (){

	// 接口组 无需授权的接口
	g:= router.Group("/api")
	{
		g.POST("/index",  Demo.Index)


	}


}