package mybatis

import (
	"database/sql"
	"github.com/xormplus/xorm"
	"github.com/zhuxiujia/GoMybatis"
)



type MysqlEngine struct{
	Engine *GoMybatis.GoMybatisEngine
}



// 通过 mybatis 新建mysql连接
//示例 "root:123456@tcp(112.131.124.103:3306)/scrapy?charset=utf8",
func NewMysqlEngine(dataSource string) (*MysqlEngine,error){

	var engine = GoMybatis.GoMybatisEngine{}.New()
	db,err := engine.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	return &MysqlEngine{
		&engine,
	},err
}


//一些初始化的设置
func initEngineer(engine *GoMybatis.GoMybatisEngine) {


}

