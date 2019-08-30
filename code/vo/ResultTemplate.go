package vo

import (
	"github.com/gin-gonic/gin"
	"log"
)

const SuccessCode string ="0000"
const FailCode string = "5000"
const FailCode_Validate string = "5001"


// 标准返回值模板
type ResultTemplate struct{
	Code string `json:"code" example:"0000" description:"返回码,0000表示正常,其他都是一场"`
	Message string `json:"message" example:"" description:"响应的消息,包含错误或正确的消息"`
	Success bool `json:"success" example:"true" description:"返回值的自我判断,true|false"`
	Data interface{} `json:"data" example:"{}" description:"返回的数据,可以是数字,字符,对象,数组等标准JSON格式"`
}

// 标准请求模板
type RequestTemplate struct {
	Token string `json:"token" text:"请求接口的令牌,一般由登录获取" required:"true"  example:"e10adc3949ba59abbe56e057f20f883e"`
	Timestamp string	`json:"timestamp" description:"当前时间戳" required:"true" `
	ClientType int	`json:"clientType" description:"客户端类型 1安卓,2IOS,3WAP,4PC,5微信小程序 6支付宝小程序 7 百度小程序 8 安卓联盟的快应用 9 待定" required:"true" `
	Sign string 	`json:"sign" description:"签名,特殊算法后的签名,用于接口非法调用判定" required:"true" `
	Data interface{} `json:"data" swaggertype:"object" example:"{}" description:"可以是数字,字符,对象,数组等标准JSON格式" required:"true" `
}






func Success(data interface{}) map[string]interface{}{
	return gin.H{
		"code":    SuccessCode,
		"message": "",
		"success": true,
		"data":    data,
	}

}

//包含全部参数的Fail
func Fail(code string,message string,data interface{})  *ResultTemplate{
	if code== SuccessCode {
		log.Fatalln("错误代码使用了SuccessCode,请修正");
		code = FailCode
	}
	return &ResultTemplate{
		Code:code,
		Message:message,
		Success:false,
		Data: data,
	}
}

// 只传递code message
func FailCodeMassage(code string,message string)  *ResultTemplate{
	if code== SuccessCode {
		log.Fatalln("错误代码使用了SuccessCode,请修正");
		code = FailCode
	}
	return &ResultTemplate{
		Code: code,
		Message:message,
		Success:false,
		Data: "",
	}
}

//只传递message
func Fail5000(message string)  *ResultTemplate{

	return &ResultTemplate{
		Code: FailCode,
		Message:message,
		Success:false,
		Data: "",
	}
}

//  传递message和data
func FailData(message string,data interface{})  *ResultTemplate{

	return &ResultTemplate{
		Code: FailCode,
		Message:message,
		Success:false,
		Data: data,
	}
}


//  传递message和data
func FailValidate(message string)  *ResultTemplate{

	return &ResultTemplate{
		Code: FailCode_Validate,
		Message: message,
		Success:false,
		Data: nil,
	}
}
