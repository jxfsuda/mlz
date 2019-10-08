package scrapyService

import (
	"crypto/tls"
	"fmt"
	"github.com/EDDYCJY/fake-useragent"
	"github.com/kirinlabs/HttpRequest"
	"golang.org/x/text/encoding/simplifiedchinese"
	"regexp"
	"strconv"
	"strings"
)
var reged_meta_charset = regexp.MustCompile("<meta .* content=\".* charset=(.*)\".*/>")
type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		var decodeBytes,_=simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str= string(decodeBytes)

	default:
		str = string(byte)
	}

	return str
}


func GetHtml(url string,param map[string]interface{}) string{

	req := HttpRequest.NewRequest().Debug(true).DisableKeepAlives(false).SetTimeout(500)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	req.SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Connection": "keep-alive",
		"User-Agent":browser.Random(),
	})
	resp,err :=req.Get(url,param)
	if err!=nil {
		panic(err.Error())
	}
	code:= resp.StatusCode()
	if code>=400{
		panic("获取列表错误,服务器响应"+ strconv.Itoa(code))
	}

	body, err := resp.Body()
	if err != nil {
		panic(err.Error())
	}
	b:= string(body)
	contentType:=resp.Response().Header.Get("content-type")
	if strings.Contains(contentType,"charset"){  //一句这个判断字符编码
		c:=strings.Split(contentType,"charset=")[1]
		if strings.Contains(c,"gb"){
			b=ConvertByte2String(body,GB18030)
		}
	}else if strings.Contains(b,"charset=") {
		ss:=reged_meta_charset.FindStringSubmatch(b)
		c:=ss[1]
		if strings.Contains(c,"gb"){
			b=ConvertByte2String(body,GB18030)
		}
	}

	fmt.Println(contentType)
	return b

}


func PostHtml(url string,param map[string]interface{}) string{

	req := HttpRequest.NewRequest().Debug(true).DisableKeepAlives(false).SetTimeout(500)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	req.SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Connection": "keep-alive",
		"User-Agent":browser.Random(),
	})
	resp,err :=req.Post(url,param)
	if err!=nil {
		panic(err.Error())
	}
	code:= resp.StatusCode()
	if code>=400{
		panic("获取列表错误,服务器响应"+ strconv.Itoa(code))
	}

	body, err := resp.Body()
	if err != nil {
		panic(err.Error())
	}

	return string(body)

}


func PostJson(url string,param map[string]interface{}) string{

	req := HttpRequest.NewRequest().Debug(true).DisableKeepAlives(false).SetTimeout(500)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	req.SetHeaders(map[string]string{
	"Content-Type": "application/json",
	"Connection": "keep-alive",
	"User-Agent":browser.Random(),
	})
	resp,err :=req.JSON().Post(url,param)
	if err!=nil {
	panic(err.Error())
	}
	code:= resp.StatusCode()
	if code>=400{
	panic("获取列表错误,服务器响应"+ strconv.Itoa(code))
	}

	body, err := resp.Body()
	if err != nil {
	panic(err.Error())
	}

	return string(body)

}