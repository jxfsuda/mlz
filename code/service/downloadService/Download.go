package downloadService

import (
	"log"
	"mlz/code/config"
	"mlz/code/vo/req"
	"mlz/iolib/httpUtils"
	"os"
)

func DownloadBook(params *req.DownloadReqVO){


	var GroupName = params.GroupName
//	ReferUrl:=params.ReferUrl
	Items:=params.Items

	var path = config.AppConfigObject.DownLoadRootPath+GroupName+"/"

	os.MkdirAll(path,os.ModePerm)
	for _,v :=range Items{

	 	log.Println(path+v.Desc)
		httpUtils.Download(v.Url,path+v.Desc, func(err error) {
			log.Println(err)
		})

	}



	//httpUtils.Download()




}
