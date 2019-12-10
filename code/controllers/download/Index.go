package download

import (
	"github.com/gin-gonic/gin"
	"log"
	"mlz/code/service/downloadService"
	"mlz/code/vo/req"
	"mlz/iolib/vo"
	"net/http"
)

// @Summary 获取数据
// @Produce  json
// @Param	body	body	req.DownloadReqVO   true        "参数对象,注意,此参数应该被包含在通用参数的data属性内"
// @Accept json
// @Success 200 {object}   req.DownloadReqVO
// @Router /api/download/PostRecord [post]
func PostRecord(c *gin.Context) {

	req:= &req.DownloadReqVO{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, vo.Fail5000(err.Error()))
		return
	}

	log.Println(req)
	downloadService.DownloadBook(req)
	c.JSON(http.StatusOK, vo.Success(nil))
}
