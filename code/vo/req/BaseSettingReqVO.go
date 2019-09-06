package req

import (
	"mlz/code/entity"
   // "mlz/iolib/time"
)

// @Description 系统设置的业务请求对象示例
type BaseSettingReqVO struct{
	entity.BaseSetting

	//下面写扩展业务字段
	// Ext string `json:"ext" description:"扩展字段说明"`

	// 基类中默认禁止这2个字段的序列化, 如果界面需要,可以在这里 覆盖设置
	//Sys_cdate time.LocalTimeSecond	`xorm:"timestamp not null 'sys_cdate'" description:"创建日期" json:"Sys_cdate" `
	//Sys_cuser int	`xorm:"int(11) not null 'sys_cuser'" description:"创建人" json:"Sys_cuser" `
}