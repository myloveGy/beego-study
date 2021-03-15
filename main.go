package main

// 引入包名
import (
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinxing-go/mysql"

	"project/connection"
	_ "project/routers"
)

// 初始化处理
func init() {
	// 加载MySQL配置文件
	mysqlConfig, _ := beego.AppConfig.GetSection("mysql")
	connection.DB = mysql.NewMySQL(&mysql.Config{
		Driver: "mysql",
		Dsn: fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True",
			mysqlConfig["dbuser"],
			mysqlConfig["dbpass"],
			mysqlConfig["dbhost"],
			mysqlConfig["dbport"],
			mysqlConfig["dbname"],
		) + "&loc=Asia%2FShanghai",
		ShowSql: true,
	})
}

// 执行主函数
func main() {
	beego.Run()
}
