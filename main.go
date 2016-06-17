package main

// 引入包名
import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "project/routers"
	"strings"
)

// 初始化处理
func init() {
	// 加载MySQL配置文件
	mysql, _ := beego.AppConfig.GetSection("mysql")
	db := fmt.Sprintf("%v:%v@/%v", mysql["dbuser"], mysql["dbpass"], mysql["dbname"])
	orm.RegisterDataBase("default", "mysql", db)
}



// 执行主函数
func main() {
	beego.Run()
}
