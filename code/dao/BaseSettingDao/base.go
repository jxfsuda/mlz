package BaseSettingDao

import (
	"mlz/code/config"
	"mlz/code/entity"
	"mlz/iolib/vo"
)

const TableName="base_setting"

func wrapAlias(tableAlias string,field string) string{
	if tableAlias=="" {
		return field
	}else{
		return tableAlias+"."+ field
	}
}
func dynamicWhereCondition(param *entity.BaseSetting,tableAlias string) string{
	var where = " 1=1 "

	 if param.Name!="" {
	 	where += " and "+ wrapAlias(tableAlias,"name") +"=?name "
	 }


	return where
}


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

func FindByPager(param *entity.BaseSetting,pager *vo.Pager) ([]entity.BaseSetting  , error){
	var entities =  make([]entity.BaseSetting,pager.PageSize)
	i,err := config.AppConfigObject.Db.Engine.SqlTemplateClient("findByPager.jet").FindAndCount(entities,param)
	if err!=nil {
		return nil , err
	}
	pager.TotalRecords=i
	return entities,nil
}