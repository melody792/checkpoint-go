package initial

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitSql() {
	user     := "pzhbao"
	password := "pzhbao"
	host     := "localhost"
	port     := 3306
	dbname   := "go_userinfo"
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	//orm.SetMaxIdleConns("default",30)
	//orm.SetMaxOpenConns("default",30)
	orm.Debug = true //可能存在性能问题，不建议使用在产品模式
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user,password, host, port, dbname))
}