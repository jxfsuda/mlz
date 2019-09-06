package xorm

// @description 分页参数
type Pager struct {
	PageSize int 	`json:"size" description:"每页展示条数"`
	PageNumber int 	`json:"page" description:"页码"`

	//以下是返回值需要的
	TotalRecord int 	`json:"total" description:"总记录数"`
}












