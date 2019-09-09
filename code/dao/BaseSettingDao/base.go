package BaseSettingDao



import (
	"mlz/code/config"
	"mlz/code/entity"
	"mlz/iolib/xorm"
)

func  GetAll() (*[]entity.BaseSetting , error){
	var entities []entity.BaseSetting
	err := config.AppConfigObject.Db.Engine.SqlMapClient("base_setting_getAll").Find(&entities)
	if err!=nil {
		return nil , err
	}
	return &entities,nil
}


func GetById(id string) (*entity.BaseSetting , error){
	var entity  entity.BaseSetting
	res := config.AppConfigObject.Db.Engine.SqlMapClient("base_setting_getById",id).GetFirst(&entity)
	if res.Error!=nil {
		return nil , res.Error
	}
	return &entity,nil
}

func GetByPager(param *entity.BaseSetting,pager xorm.Pager) (*[]entity.BaseSetting  , error){
	var entities []entity.BaseSetting
	err := config.AppConfigObject.Db.Engine.SqlMapClient("base_setting_getById",param).Find(&entities)
	if err!=nil {
		return nil , err
	}
	return &entities,nil
}

