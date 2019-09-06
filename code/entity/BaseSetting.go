package entity

import (
	"mlz/iolib/time"
)

// @Description 系统设置
type BaseSetting struct {
	Id string	`xorm:"char(32)  'id'" description:"记录id" json:"id" `
	Name string	`xorm:"varchar(50) not null 'name'" description:"设置名称" json:"name" `
	Keyer string	`xorm:"varchar(50)  'keyer'" description:"key,代码定义" json:"keyer" `
	Val string	`xorm:"varchar(200) not null 'val'" description:"设置值" json:"val" `
	Remark string	`xorm:"varchar(500) not null 'remark'" description:"备注" json:"remark" `
	Status int	`xorm:"int(11) not null 'status'" description:"状态" json:"status" `
	Sys_cdate time.LocalTimeSecond	`xorm:"datetime not null 'sys_cdate'" description:"创建日期" json:"-" `
	Sys_cuser int	`xorm:"int(11) not null 'sys_cuser'" description:"创建人" json:"-" `
}




