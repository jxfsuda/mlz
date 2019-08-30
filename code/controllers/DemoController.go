package controllers

import (
	"mlz/code/vo"
	"net/http"
)
import "github.com/gin-gonic/gin"


type Demo struct{
	Id int `json:"id" example:"1" description:"编号"`
	Name string `json:"name" example:"jif" description:"名称"`
}



type DemoController struct {
	*vo.RequestTemplate	`description:"继承标准结构"`
	Data Demo 	`json:"data" description:"业务数据" "反倒是考虑; "`
}




// @Summary 演示
// @Produce  json
// @Param	body	body	controllers.DemoController   true        "参数对象,注意,此参数应该被包含在通用参数的data属性内"
// @Accept json
// @Success 200 {string} json "{"code":0000,"data":{},"message":"","success":true}"
// @Router /api/v1/demo/index [post]
func (a *DemoController) Index(c *gin.Context) {

	var json DemoController
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vo.Success(json.Data))
}
