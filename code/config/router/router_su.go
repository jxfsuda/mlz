package router

import (
	"mlz/code/controllers/su/BaseSetting"
)

func route_su (){


	//接口组1  可以添加授权校验的 版本1的接口 , 支持内部继续分组
	g1:= router.Group("/api/su/v1")
	{

		demo:=g1.Group("/base_setting")
		{
			demo.POST("/getAll", BaseSetting.GetAll)
			demo.POST("/listData", BaseSetting.ListData)
		}




	}

}