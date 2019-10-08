package BaseSettingService

import (
	"mlz/code/dao/BaseSettingDao"
	"mlz/code/entity"
	"mlz/iolib/vo"
)
// 此文件由代码生成,请勿手工修改

func GetAll() (*[]entity.BaseSetting,error) {

	res,err:= BaseSettingDao.GetAll()
	if err!=nil {
		panic(err.Error())
	}
	return res,nil
}



func FindByPager(entity *entity.BaseSetting,pager *vo.Pager) ([]entity.BaseSetting,error) {

	res,err:= BaseSettingDao.FindByPager(entity,pager)
	if err!=nil {
		panic(err.Error())
	}
	return res,nil
}

