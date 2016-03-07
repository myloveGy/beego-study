package main

import (
	_ "app/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 初始化处理函数
func init() {
	// 加载MySQL配置文件
	mysql, _ := beego.AppConfig.GetSection("mysql")
	db := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", mysql["dbuser"], mysql["dbpass"], mysql["dbhost"], mysql["dbport"], mysql["dbname"])
	orm.RegisterDataBase("default", "mysql", db)
}

func main() {
	beego.Run()
}
