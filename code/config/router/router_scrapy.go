package router

import (
	"mlz/code/controllers/download"
	"mlz/code/controllers/scrapy"
)

func route_scrapy (){


	//接口组1  可以添加授权校验的 版本1的接口 , 支持内部继续分组
	g1:= router.Group("/api/scrapy/")
	{

		demo:=g1.Group("/entry")
		{
			demo.POST("/GetList", scrapy.GetList)
			demo.POST("/GetData", scrapy.GetData)

		}
	}


	g2:= router.Group("/api/download/")
	{


		g2.POST("/PostRecord", download.PostRecord)


	}





}