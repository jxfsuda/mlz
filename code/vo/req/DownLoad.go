package req

// 远程下载请求结构
type DownloadReqVO struct{
	GroupName string	`json:"groupName",description:"下载组名称,比如同一书籍名,本次下载将全部追加到此目录"`
	Items	  []DownloadItem `json:"items",description:"要下载的列表"`
	ReferUrl  string `json:"referUrl",description:"来源网页,可为空"`
}

type DownloadItem struct{
	Title string `json:"title",description:"下载链文本,作为文件保存名称"`
	Url string 	`json:"url",description:"下载链接"`
	Desc  string `json:"desc",description:"下载说明"`
}