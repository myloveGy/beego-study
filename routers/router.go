package routers

import (
	"app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 固定路由
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.Router("/Admin", &controllers.AdminController{})

	// 自由路由
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.MenusController{})
	beego.AutoRouter(&controllers.OtherController{})
}
