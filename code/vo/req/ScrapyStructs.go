package req

type Link struct{
	Title string `json:"title"`
	Url string `json:"url"`
}

type GetListReq struct{
	Url string	`json:"url"`
}

type GetListRsp struct{
	Total int	`json:"total" description:"总数"`
	Links []Link	`json:"links" description:"分集链接"`
	Name string 	`json:"name" description:"书名"`
}

type DataVo struct{

	Title string `json:"title"`
	Url string `json:"url"`
	AudioUrl string `json:"audio_url"`
	Size int `json:"size"`
	Message string `json:"message"`

}