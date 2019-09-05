package xorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
	"log"
)




type MysqlEngine struct{
	Engine *xorm.Engine
}

// 新建mysql连接
//示例 "root:123456@tcp(112.131.124.103:3306)/scrapy?charset=utf8",
func NewMysqlEngine(dataSource string) (*MysqlEngine,error){

	var eng = MysqlEngine{}
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Fatalln("数据库连接失败" + err.Error())
		return nil , err
	}

	err = engine.Ping()  //测试数据库是否可用
	if err != nil {
		log.Fatalln("数据库ping失败" + err.Error())
		return nil , err
	}

//xorm reverse mysql 'root:442503@tcp(111.231.104.83:3306)/scrapy?charset=utf8' conf/mapper/mysql
	initEngineer(engine)

	eng.Engine = engine


	return &eng,err
}


//一些初始化的设置
func initEngineer(engine *xorm.Engine) {
	engine.SetTableMapper(core.SnakeMapper{})  //表名采用驼峰
	engine.SetColumnMapper(core.SameMapper{})  //字段为same,即数据库是什么,映射就是什么 映射规则 见 https://www.kancloud.cn/xormplus/xorm/167084

	//日志打印SQL
	engine.ShowSQL(true)

	//设置连接池的空闲数大小
	engine.SetMaxIdleConns(2)
	//设置最大打开连接数
	engine.SetMaxOpenConns(10)






	//注册映射文件
	err := engine.RegisterSqlMap(xorm.Xml("./conf/mapper/mysql", ".xml"))
	if err != nil {
		log.Fatal(err)
	}

	// xorm  reverse mysql  "root:442503@tcp(111.231.104.83:3306)/scrapy?charset=utf8" /data/go-workspace/mlz/conf/mapper/mysql
	//注册sql模板  jet 直接写 > < = 而不是 gt lt , 且用 {{if data.count>1 }}  ... {{else}}} ... {{end}} 个人喜好, 都差不多,看的顺眼而已
	err = engine.RegisterSqlTemplate(xorm.Default("./conf/mapper/mysql", ".jet"))
	if err != nil {
		log.Fatal(err)
	}

	//开启SqlMap配置文件和SqlTemplate配置文件更新监控功能，将配置文件更新内容实时更新到内存，如无需要可以不调用该方法
	err = engine.StartFSWatcher()
	if err != nil {
		log.Fatal(err)
	}
}



