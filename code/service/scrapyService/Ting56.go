package scrapyService

import (
	"github.com/PuerkitoBio/goquery"
	"mlz/code/vo/req"
	"strings"
)
func ProcessTing56(url string)  *req.GetListRsp{

	body:=GetHtml(url,nil)
 	r:=strings.NewReader(body)
	doc,err:=goquery.NewDocumentFromReader(r)
	if err!=nil {
		panic(err)
	}
	var ret = &req.GetListRsp{
		Total:0,
		Name:"",
	}
	name := doc.Find(".content .left .tit h1").Text();
	 as := doc.Find("#vlink_1 ul li a")
	len := as.Length()
	ret.Total = len
	ret.Name=strings.ReplaceAll(name,"有声小说","")

	links:= make([]req.Link,len)

	as.Each(func(i int, selection *goquery.Selection) {
		t:=strings.TrimSpace(selection.Text())
		href,_:=selection.Attr("href")

		links[i]= req.Link{
			Title:t,
			Url: "http://www.ting56.com"+href ,
		}
	})
	ret.Links = links

	return ret
}


func ProcessTing56Data(url string)