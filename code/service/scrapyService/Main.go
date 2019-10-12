package scrapyService

import (
	"mlz/code/vo/req"
	"regexp"
	"strings"
	"sync"
)

var regex = regexp.MustCompile("://([\\w\\.]{1,})/")

var wg sync.WaitGroup

var name_prefix string

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
		ps := strings.LastIndex(v.Title,"第")

		p:=v.Title[:ps]

		if p!="" {
			for key, value := range Zh_num_2_num {
				p=strings.ReplaceAll(p,key,value)
			}
			name_prefix=p
		}
		link:=v.Url

		go func(link string,p string) {
			GetOneData(link, p)
		}(link,name_prefix)

	}
	wg.Wait()
	return nil
}


func GetOneData(link string,p string) interface{}{
	wg.Add(1)
	defer wg.Done()
	mts:=regex.FindStringSubmatch(link)
	var host = mts[1]
	switch host{
	case "www.ting56.com", "ting56.com":
		return ProcessTing56Data(link,p)
		break
	default:
		panic("暂不支持此站点")
	}

	return nil
}
