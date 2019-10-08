package scrapyService

import (

	"mlz/code/vo/req"
	"regexp"
)

var regex = regexp.MustCompile("://([\\w\\.]{1,})/")




// 开始处理
func Process(reqVO *req.GetListReq) *req.GetListRsp{

	book_url := reqVO.Url

	mts:=regex.FindStringSubmatch(book_url)
	var host = mts[1]

	switch host{
		case "www.ting56.com", "ting56.com":
			return ProcessTing56(book_url)
			break
	default:
		panic("暂不支持此站点")
	}

	return nil
}


func GetData(reqVO *req.GetListRsp) *req.GetListRsp{

	links := reqVO.Links
	for _,v:=range links{



	}

	return nil
}


func GetOneData(link *req.Link){
	mts:=regex.FindStringSubmatch(link.Url)
	var host = mts[1]
	switch host{
	case "www.ting56.com", "ting56.com":
		return ProcessTing56Data(book_url)
		break
	default:
		panic("暂不支持此站点")
	}


}
