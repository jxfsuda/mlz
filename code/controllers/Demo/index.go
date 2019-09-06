package Demo

import (
	"github.com/gin-gonic/gin"
	"mlz/code/service/BaseSettingService"
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


	//var json entity.BaseSetting
	//if err := c.ShouldBindJSON(&json); err != nil {
	//
	//	return
	//}
	BaseSettingService.GetAll()

	c.JSON(http.StatusOK, vo2.Success(""))
}

