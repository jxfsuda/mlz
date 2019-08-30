package vo

//@Description 标准返回值模板
type ResultTemplate struct{
	Code string `json:"code" example:"0000" description:"返回码,0000表示正常,其他都是一场"`
	Message string `json:"message" example:"" description:"响应的消息,包含错误或正确的消息"`
	Success bool `json:"success" example:"true" description:"返回值的自我判断,true|false"`
	Data interface{} `json:"data" example:"{}" description:"返回的数据,可以是数字,字符,对象,数组等标准JSON格式"`
}

//@Description 标准请求模板
type RequestTemplate struct {
	Token string `json:"token" description:"请求接口的令牌,一般由登录获取" required:"true" `
	Timestamp string	`json:"" description:"" required:"true" `
	ClientType int	`json:"" description:"客户端类型 1安卓,2IOS,3WAP,4PC,5微信小程序 6支付宝小程序 7 百度小程序 8 安卓联盟的快应用 9 待定" required:"true" `
	Sign string 	`json:"" description:"签名,特殊算法后的签名,用于接口非法调用判定" required:"true" `
	Data interface{} `json:"data" swaggertype:"object" example:"{}" description:"可以是数字,字符,对象,数组等标准JSON格式" required:"true" `
}

ResultTemplate :=ResultTemplate{}

func success(data interface{}){

}