package BaseSetting


import (
	"github.com/gin-gonic/gin"
	"mlz/code/service/BaseSettingService"
	"mlz/iolib/vo"
	"net/http"
)


// @Summary 列表数据
// @Produce  json
// @Param	body	body	req.BaseSettingReqVO   true        "参数对象,注意,此参数应该被包含在通用参数的data属性内"
// @Accept json
// @Success 200 {string} json "{"code":0000,"data":{},"message":"","success":true}"
// @Router /api/su/v1/base_setting/GetAll [post]
func GetAll(c *gin.Context) {
	//if err := c.ShouldBindJSON(&json); err != nil {
	//
	//	return
	//}
	//var param = &entity.BaseSetting{}
	res, _ := BaseSettingService.GetAll()

	c.JSON(http.StatusOK, vo.Success(res))
}


func ListData(c *gin.Context) {
	//if err := c.ShouldBindJSON(&json); err != nil {
	//
	//	return
	//}
	//var param = &entity.BaseSetting{}
	res, _ := BaseSettingService.GetAll()

	c.JSON(http.StatusOK, vo.Success(res))
}
