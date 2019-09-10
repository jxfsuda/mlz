package BaseSettingService

import (
	"mlz/code/dao/BaseSettingDao"
	"mlz/code/entity"
	"mlz/iolib/xorm"
)
//数据库操作引擎, 默认的db,如果有其他db,自行创建即可

func GetAll() (*[]entity.BaseSetting,error) {

	 res,err:= BaseSettingDao.GetAll()
	 if err!=nil {
	 	panic(err.Error())
	 }
	 return res,nil
}



func FindByPager(entity *entity.BaseSetting,pager *xorm.Pager) (*[]entity.BaseSetting,error) {

	res,err:= BaseSettingDao.FindByPager(entity,pager)
	if err!=nil {
		panic(err.Error())
	}
	return res,nil
}



