package entity

import (
	_ "github.com/go-sql-driver/mysql"
	"mlz/iolib/timeUtils"
)

// @Description 系统设置
type BaseSetting struct {
	Id        string                    `xorm:"char(32)  'id'" description:"记录id" json:"id" `
	Name      string                    `xorm:"varchar(50) not null 'name'" description:"设置名称" json:"name" `
	Keyer     string                    `xorm:"varchar(50)  'keyer'" description:"key,代码定义" json:"keyer" `
	Val       string                    `xorm:"varchar(200) not null 'val'" description:"设置值" json:"val" `
	Remark    string                    `xorm:"varchar(500) not null 'remark'" description:"备注" json:"remark" `
	Status    int                       `xorm:"int(11) not null 'status'" description:"状态" json:"status" `
	Sys_cdate timeUtils.LocalTimeSecond `xorm:"datetime not null 'sys_cdate'" description:"创建日期" json:"-" `
	Sys_cuser int                       `xorm:"int(11) not null 'sys_cuser'" description:"创建人" json:"-" `
}



type BaseSettingMapper struct {
	//调用即可生成sql
	SelectTemplete      func(name string) ([]BaseSetting, error)
	InsertTemplete      func(arg *BaseSetting) (int64, error)
	InsertTempleteBatch func(args []BaseSetting) (int64, error) `mapperParams:"args"`
	//生成sql(带有乐观锁.逻辑删除)  update base_setting set name = #{name},remark=#{remark},version=#{version+1} where delete_flag = 1 and id = #{id} and version = #{version}
	UpdateTemplete      func(arg BaseSetting) (int64, error) `
	DeleteTemplete      func(id string) (int64, error)     `mapperParams:"id"`
}

