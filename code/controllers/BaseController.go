package controllers

import (
	"github.com/gin-gonic/gin"
	"mlz/code/vo"
	"net/http"
)

type BaseController struct {

}



func (c *gin.Context) ErrorResponse(msg string) {
	c.JSON(http.StatusOK, vo.FailValidate(msg))
}
