package entity

import "time"

// @Summary 系统设置
type BaseSetting struct {
	Id string	`xorm:"char(32)  'id'" description:"记录id"`
	Name string	`xorm:"varchar(50) not null 'name'" description:"设置名称"`
	Keyer string	`xorm:"varchar(50)  'keyer'" description:"key,代码定义"`
	Val string	`xorm:"varchar(200) not null 'val'" description:"设置值"`
	Remark string	`xorm:"varchar(500) not null 'remark'" description:"备注"`
	Status int	`xorm:"int(11) not null 'status'" description:"状态"`
	Sys_cdate time.Time	`xorm:"datetime not null 'sys_cdate'" description:"创建日期"`
	Sys_cuser int	`xorm:"int(11) not null 'sys_cuser'" description:"创建人"`
}

