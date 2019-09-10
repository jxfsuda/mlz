package BaseSettingDao



import (
	"mlz/code/config"
	"mlz/code/entity"
	"mlz/iolib/xorm"
)

const TableName="base_setting"



func  GetAll() (*[]entity.BaseSetting , error){
	var entities []entity.BaseSetting
	err := config.AppConfigObject.Db.Engine.SqlMapClient(TableName+"_getAll").Find(&entities)
	if err!=nil {
		return nil , err
	}
	return &entities,nil
}


func GetById(id string) (*entity.BaseSetting , error){
	var entity  entity.BaseSetting
	res := config.AppConfigObject.Db.Engine.SqlMapClient(TableName+"_getById",id).GetFirst(&entity)
	if res.Error!=nil {
		return nil , res.Error
	}
	return &entity,nil
}

func FindByPager(param *entity.BaseSetting,pager *xorm.Pager) (*[]entity.BaseSetting  , error){
	var entities []entity.BaseSetting



	err := config.AppConfigObject.Db.Engine.SqlMapClient(TableName+"_findByPager",param).Limit(pager.PageNumber,pager.PageSize).Find(&entities)
	if err!=nil {
		return nil , err
	}
	return &entities,nil
}

