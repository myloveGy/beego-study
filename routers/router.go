package routers

import (
	"project/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 前台显示
	beego.Router("/", &controllers.IndexHomeController{})
	beego.Router("/ajax", &controllers.IndexHomeController{}, "*:Ajax")
	beego.Router("/login", &controllers.IndexHomeController{}, "*:Login")
	beego.Router("/logout", &controllers.IndexHomeController{}, "*:Logout")
	beego.Router("/image", &controllers.IndexHomeController{}, "*:Image")
	beego.Router("/insert", &controllers.IndexHomeController{}, "*:Insert")
	beego.Router("/add", &controllers.IndexHomeController{}, "*:Add")
	beego.Router("/upload", &controllers.IndexHomeController{}, "*:Upload")
	beego.Router("/article", &controllers.ArticleHomeController{}, "*:Index")
	beego.Router("/article/list", &controllers.ArticleHomeController{}, "*:List")
	beego.Router("/article/view/:id([0-9]+)", &controllers.ArticleHomeController{}, "*:View")

	// 后台显示
	beego.Router("/admin", &controllers.SiteController{}, "*:Index")
	beego.Router("/admin/login", &controllers.SiteController{}, "*:Login")
	beego.Router("/admin/logout", &controllers.SiteController{}, "*:Logout")
	beego.Router("/admin/site", &controllers.AdminAdminController{}, "*:Site")
	beego.Router("/admin/update", &controllers.AdminAdminController{}, "*:Update")
	beego.Router("/admin/upload", &controllers.AdminAdminController{}, "*:Upload")
}
