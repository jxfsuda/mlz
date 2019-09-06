package BaseSettingService

import (
	"mlz/code/dao/BaseSettingDao"
	"mlz/code/entity"
)
//数据库操作引擎, 默认的db,如果有其他db,自行创建即可

func GetAll() (*[]entity.BaseSetting,error) {

	 res,err:= BaseSettingDao.GetAll()
	 if err!=nil {
	 	panic(err.Error())
	 }
	panic("测试recovery")
	 return res,nil
}