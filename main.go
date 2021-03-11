package main

// 引入包名
import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	_ "project/routers"
)

// 初始化处理
func init() {
	// 加载MySQL配置文件
	mysql, _ := beego.AppConfig.GetSection("mysql")
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8",
		mysql["dbuser"],
		mysql["dbpass"],
		mysql["dbhost"],
		mysql["dbport"],
		mysql["dbname"],
	)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.Debug = true
}

// 执行主函数
func main() {
	beego.Run()
}
