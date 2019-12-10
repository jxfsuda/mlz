package scrapyService

import (
	"mlz/code/vo/req"
	"regexp"
	"sync"
)

var regex = regexp.MustCompile("://([\\w\\.]{1,})/")

var wg sync.WaitGroup


// 开始处理
func Process(reqVO *req.GetListReq) *req.GetListRsp{

	book_url := reqVO.Url

	mts:=regex.FindStringSubmatch(book_url)
	var host = mts[1]

	switch host{
		case "www.ting56.com", "ting56.com":
			return ProcessTing56(book_url)
			break
	case "www.ting89.com", "ting89.com":
		return ProcessTing89(book_url)
		break
	default:
		panic("暂不支持此站点")
	}

	return nil
}


func GetData(reqVO *req.GetListRsp) *req.GetListRsp{

	links := reqVO.Links
	for idx,v:=range links{
		link:=v.Url
		go func(link string,p string,idx int) {
			GetOneData(link, p,idx)
		}(link,reqVO.Name,idx)

	}
	wg.Wait()
	return nil
}


func GetOneData(link string,p string, idx int) interface{}{
	wg.Add(1)
	defer wg.Done()
	mts:=regex.FindStringSubmatch(link)
	var host = mts[1]
	switch host{
	case "www.ting56.com", "ting56.com":
		return ProcessTing56Data(link,p,idx)
		break
	case "www.ting89.com", "ting89.com":
		return ProcessTing89Data(link,p,idx)
		break
	default:
		panic("暂不支持此站点")
	}

	return nil
}
