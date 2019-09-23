package Demo

import (
	"github.com/gin-gonic/gin"
	"mlz/code/entity"
	"mlz/code/service/BaseSettingService"
	"mlz/iolib/mybatis"
	vo2 "mlz/iolib/vo"
	"net/http"
)

// @Summary 演示
// @Produce  json
// @Param	body	body	req.BaseSettingReqVO   true        "参数对象,注意,此参数应该被包含在通用参数的data属性内"
// @Accept json
// @Success 200 {string} json "{"code":0000,"data":{},"message":"","success":true}"
// @Router /api/v1/demo/index [post]
func Index(c *gin.Context) {


	 var vo = new(entity.BaseSetting)
	//if err := c.ShouldBindJSON(&json); err != nil {
	//
	//	return
	//}
	BaseSettingService.FindByPager(vo,&mybatis.Pager{})

	c.JSON(http.StatusOK, vo2.Success(""))
}

