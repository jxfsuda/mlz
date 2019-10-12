package scrapy

import (
	"github.com/gin-gonic/gin"
	"mlz/code/service/scrapyService"
	req2 "mlz/code/vo/req"
	"mlz/iolib/vo"
	"net/http"
)

// @Summary 获取目录
// @Produce  json
// @Param	body	body	req.GetListReq   true        "参数对象,注意,此参数应该被包含在通用参数的data属性内"
// @Accept json
// @Success 200 {string} json "{"code":0000,"data":{},"message":"","success":true}"
// @Router /api/scrapy/entry/GetList [post]
func GetList(c *gin.Context) {

	req:= &req2.GetListReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, vo.Fail5000(err.Error()))
		return
	}
	ret:=scrapyService.Process(req)

	c.JSON(http.StatusOK, vo.Success(ret))
}

// @Summary 获取数据
// @Produce  json
// @Param	body	body	req.GetListRsp   true        "参数对象,注意,此参数应该被包含在通用参数的data属性内"
// @Accept json
// @Success 200 {string} json "{"code":0000,"data":{},"message":"","success":true}"
// @Router /api/scrapy/entry/GetData [post]
func GetData(c *gin.Context) {

	req:= &req2.GetListRsp{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, vo.Fail5000(err.Error()))
		return
	}

	ret:=scrapyService.GetData(req)

	c.JSON(http.StatusOK, vo.Success(ret))
}
