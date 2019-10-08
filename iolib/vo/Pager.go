package vo

// @description 分页参数
type Pager struct {
	PageSize int 	`json:"size" description:"每页展示条数"`
	PageNumber int 	`json:"page" description:"页码"`

	//以下是返回值需要的
	TotalRecords int64 	`json:"total" description:"总记录数"`
}


func (c *Pager) Limit() int{
	return (c.PageNumber-1)*c.PageSize
}


func  (c *Pager) Offset() int{
	return c.PageSize
}

func  (c *Pager) Page() int{
	if c.PageNumber < 1 {
		return 1
	}else{
		return  c.PageNumber
	}
}











