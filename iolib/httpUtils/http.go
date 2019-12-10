package httpUtils

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func NewClient() *http.Client{
	client := http.Client{
		//涵盖了用于重定向花费的时间
		Timeout: 900000 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 10,
			DisableKeepAlives:false,
			// TLS握手使用的时间 网速较慢,或fq站点,不设置,将直接导致超时发生
			TLSHandshakeTimeout: 1000000 * time.Second,
			//读取响应报文头使用的时间
			ResponseHeaderTimeout: 1011000 * time.Second,
			//客户端在发送一个包含：100-continue的http报文头后，等待收到一个go-ahead响应报文所用的时间
			ExpectContinueTimeout: 111100 * time.Second,
		},
	}
	return &client
}



func GetResponse(url string,fn func(interface{} ,error)){
	client:=NewClient()
	req, err := http.NewRequest("GET", url,nil)
	if err != nil {
		return
	}
	req.Close = true //阻止连接被重用，可以有效的防止这个问题，也就是Http的短连接
	req.Header.Add("Connection", "close")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err!=nil{
		panic(err)
	}


	bt,err:=ioutil.ReadAll(resp.Body)
	defer fn(string(bt),err)
}




//下载文件
//基于go 1.12.4
func  Download(url string,filePath string,fn func(err error)){
	fi, err := os.Stat(filePath)
	if err==nil {  //文件已经存在

		if fi.Size()>624000 {
			log.Println("文件已存在,文件大小:",fi.Size()/(1024*1024),"M")
			return
		}
	}
	client:=NewClient()
	req, err := http.NewRequest("GET", url,nil)
	if err != nil {
		return
	}
	req.Close = true //阻止连接被重用，可以有效的防止这个问题，也就是Http的短连接
	req.Header.Add("Connection", "close")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299")

	resp, err := client.Do(req)
	if err!=nil{
		panic(err)
		return
	}

	defer resp.Body.Close()

	//stat:=resp.Status
	stat:=resp.StatusCode
	if stat>=400 {
		log.Println("下载文件错误,响应码"+resp.Status , url)
		err = errors.New("远程文件未找到"+ url)
		return
	}
	//fmt.Printf("%t,%v",stat,stat)
	if err!=nil{
		panic( err)
		return
	}

	_, err  = os.Create(filePath)
	if err!=nil{
		panic( err)
		return
	}
	bt,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		panic(err)
		return
	}
	err=ioutil.WriteFile(filePath,bt,0644)
	if err!=nil{
		panic(err)
		return
	}
	defer fn(err)
}