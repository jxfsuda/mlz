package scrapyService

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"mlz/code/vo/req"
	"os"
	"regexp"
	"strconv"
	"strings"
)




var regex_js = regexp.MustCompile("FonHen_JieMa\\('(.*)'\\).split")

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


func ProcessTing56Data(url string,title string)  *req.GetListRsp{

	body:=GetHtml(url,nil)
	r:=strings.NewReader(body)
	doc,err:=goquery.NewDocumentFromReader(r)
	if err!=nil {
		panic(err)
	}
	doc.Find("script").Each(func(i int, selection *goquery.Selection) {

		src,_:=selection.Attr("src")
		if src!="" {
			return
		}
		js:=selection.Text()
		if !strings.Contains(js,"FonHen_JieMa(") {
			return
		}

		matchs:=regex_js.FindStringSubmatch(js)
		if len(matchs) <2{
			return
		}
		str:= matchs[1]

		strs:= FonHen_JieMa(str)

		u:=strs[0]
	//	c:=strs[1]
		t:=strs[2]
		log.Println("--->" + u)
		filePath:="/data/txt/"
		if t=="tc" {
			// 再次请求
			splits:=strings.Split(u,"/")
			u=splits[0]+"/"+splits[1]+"/play_"+ splits[1]+"_"+splits[2]+".htm"
			body:=GetHtml("http://www.ting56.com/player/tingchina.php?url="+u,nil)
			res := &req.TingChinaResult{}
			_ = json.Unmarshal(([]byte)(body),res)
			u=res.Url
		}else{
			u=strings.ReplaceAll(u,":82","")
		}
		if u==""{
			log.Println(url+" 没有找到音频地址: "+ u)
		}
		up:=strings.Split(u,"?")[0]
		us:=strings.Split(up,"/")
		name:= us[len(us)-1]
		//直接下载
		file:=filePath+title+name
		fi,err:=os.Stat(file)
		if err==nil {
			if fi.Size()>1024{
				return
			}else{
				os.Remove(file)
			}
		}

		log.Println("请求URL"+ url+ " --> "+ u)
		Download(u,file)
		for i:=0;i<10;i++ {  //最多重试10次
			fi, err = os.Stat(file)
			if err == nil {
				if fi.Size() < 1024 {
					os.Remove(file)
					Download(u, file)
				}else{
					return
				}
			}
		}
	})


	return nil
}


func FonHen_JieMa(str string) []string{
	str=strings.TrimSpace(str)
	tArr:=strings.Split(str,"*")

	n:=len(tArr)

	x:=""

	for   i:=0;i<n ;i++  {
		if tArr[i]=="" {
			continue
		}
		nc,_:=strconv.Atoi(tArr[i])
		x+=string(nc)
	}

	return strings.Split(x,"&")
}