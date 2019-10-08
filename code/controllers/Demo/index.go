package Demo

import (
	"github.com/gin-gonic/gin"
	"mlz/code/entity"
	"mlz/code/service/BaseSettingService"
	"mlz/iolib/vo"
	"net/http"
)

// @Summary 演示
// @Produce  json
// @Param	body	body	req.BaseSettingReqVO   true        "参数对象,注意,此参数应该被包含在通用参数的data属性内"
// @Accept json
// @Success 200 {string} json "{"code":0000,"data":{},"message":"","success":true}"
// @Router /api/demo/v1/index [post]
func Index(c *gin.Context) {



	//if err := c.ShouldBindJSON(&json); err != nil {
	//
	//	return
	//}
	var param = &entity.BaseSetting{}
	var pager = new(vo.Pager)
	pager.PageNumber=1
	pager.PageSize=20
	res, _ := BaseSettingService.FindByPager(param,pager)

	c.JSON(http.StatusOK, vo.Success(res))
}

