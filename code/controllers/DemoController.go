package controllers

import (
	"mlz/code/vo"
	"net/http"
)
import "github.com/gin-gonic/gin"


type Demo struct{
	Id int `json:"id" example:"1" description:"编号"`
	//自定义验证,见 validator.DemoValidator DemoName
	Name string `json:"name" example:"jif" description:"名称"  binding:"required,DemoName"`
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
		c.JSON(http.StatusBadRequest, vo.FailValidate(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success(json.Data))
}

func (a *DemoController) Index1(c *gin.Context) {
	name := c.Param("name")  //获取路径参数
	firstname := c.DefaultQuery("firstname", "Guest")  //GET
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值
// // 单文件
	//        file, _ := c.FormFile("file")
	//        log.Println(file.Filename)
	//
	//        // 上传文件到指定的路径
	//        // c.SaveUploadedFile(file, dst)


	// // 多文件
	//        form, _ := c.MultipartForm()
	//        files := form.File["upload[]"]
	//
	//        for _, file := range files {
	//            log.Println(file.Filename)
	//
	//            // 上传文件到指定的路径
	//            // c.SaveUploadedFile(file, dst)
	//        }
	c.JSON(http.StatusOK, vo.Success(message+nick+name+ firstname))
}
