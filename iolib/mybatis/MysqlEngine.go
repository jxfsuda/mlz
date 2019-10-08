package mybatis
//
//import (
//	"database/sql"
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/zhuxiujia/GoMybatis"
//)
//
//
//
//type MysqlEngine struct{
//	Db *sql.DB
//	Engine *GoMybatis.GoMybatisEngine
//}
//
//
//
//// 通过 mybatis 新建mysql连接
////示例 "root:123456@tcp(112.131.124.103:3306)/scrapy?charset=utf8",
//func NewMysqlEngine(dataSource string) (*MysqlEngine,error){
//
//	var engine = GoMybatis.GoMybatisEngine{}.New()
//	db,err := engine.Open("mysql", dataSource)
//	if err != nil {
//		panic(err)
//	}
//	return &MysqlEngine{
//		db,
//		&engine,
//	},err
//}
//
//
////一些初始化的设置
//func initEngineer(engine *GoMybatis.GoMybatisEngine) {
//
//
//}
//
