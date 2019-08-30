package controllers

import (
	"mlz/code/vo"
	"net/http"
)
import "github.com/gin-gonic/gin"


type Demo struct{
	Id int `json:"id" example:"1"`
	Name string `json:"name" example:"jif"`
}


type DemoController struct{}


// @Summary 演示
// @Produce  json
// @Param	body	body	vo.RequestTemplate   true        "The object content"
// @Accept json
// @Success 200 {string} json "{"code":0000,"data":{},"message":"","success":true}"
// @Router /api/v1/demo/index [post]
func (a *DemoController) Index(c *gin.Context) {

	var json vo.RequestTemplate
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": json.Data})
}
