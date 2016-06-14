package routers

import (
	"project/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 前台显示
	beego.Router("/", &controllers.IndexController{})
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.ArticleController{})

	// 后台显示
	beego.Router("/admin", &controllers.SiteController{}, "*:Index")

	// 使用命名空间
	ns := beego.NewNamespace("/admin",
		beego.NSAutoRouter(&controllers.SiteController{}),
		beego.NSAutoRouter(&controllers.AdminController{}),
		beego.NSAutoRouter(&controllers.MenuController{}),
		beego.NSAutoRouter(&controllers.CategoryController{}),
	)

	beego.AddNamespace(ns)
}
