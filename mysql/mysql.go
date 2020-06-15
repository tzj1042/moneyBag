package mysql

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
)

var Engine *xorm.Engine

func init() {
	//mysql初始化
	var err error
	Engine, err = xorm.NewEngine("mysql", beego.AppConfig.String("mysql"))
	if err != nil {
		logs.Error("数据库连接错误:", err)
		panic(err)
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")
	Engine.SetTableMapper(tbMapper)
	//日志打印SQL
	Engine.ShowSQL(true)

	//设置连接池的空闲数大小
	Engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	Engine.SetMaxOpenConns(5)
}
