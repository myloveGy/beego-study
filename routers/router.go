package routers

import (
	"project/controllers"
	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/context"
)

func init() {
//	// 前台显示
//	beego.Router("/", &controllers.IndexHomeController{})
//	beego.Router("/ajax", &controllers.IndexHomeController{}, "*:Ajax")
//	beego.Router("/login", &controllers.IndexHomeController{}, "*:Login")
//	beego.Router("/logout", &controllers.IndexHomeController{}, "*:Logout")
//	beego.Router("/image", &controllers.IndexHomeController{}, "*:Image")
//	beego.Router("/insert", &controllers.IndexHomeController{}, "*:Insert")
//	beego.Router("/add", &controllers.IndexHomeController{}, "*:Add")
//	beego.Router("/upload", &controllers.IndexHomeController{}, "*:Upload")
//	beego.Router("/article", &controllers.ArticleHomeController{}, "*:Index")
//	beego.Router("/article/list", &controllers.ArticleHomeController{}, "*:List")
//	beego.Router("/article/view/:id([0-9]+)", &controllers.ArticleHomeController{}, "*:View")

//	beego.AutoRouter(&controllers.CategoryController{})
//	beego.AutoRouter(&controllers.MenuAdminController{})
	ns := beego.NewNamespace("/admin",
//		beego.NSCond(func(ctx *context.Context) bool {
//			if ctx.Input.Domain() == "go.com" || ctx.Input.Domain() == "localhost:8081" {
//				return true
//			}
//			return false
//		}),
		beego.NSAutoRouter(&controllers.AdminAdminController{}),
		beego.NSAutoRouter(&controllers.MenuAdminController{}),
		beego.NSAutoRouter(&controllers.CategoryController{}),
	)

	beego.AddNamespace(ns)

	// 后台显示
//	beego.Router("/admin", &controllers.SiteController{}, "*:Index")
//	beego.Router("/admin/login", &controllers.SiteController{}, "*:Login")
//	beego.Router("/admin/logout", &controllers.SiteController{}, "*:Logout")
//	beego.Router("/admin/site", &controllers.AdminAdminController{}, "*:Site")
//	beego.Router("/admin/update", &controllers.AdminAdminController{}, "*:Update")
//	beego.Router("/admin/upload", &controllers.AdminAdminController{}, "*:Upload")
//	beego.Router("/admin/menu/", &controllers.MenuAdminController{}, "*:Index")
//	beego.Router("/admin/menu/search", &controllers.MenuAdminController{}, "*:Search")
//	beego.Router("/admin/menu/update", &controllers.MenuAdminController{}, "*:Update")
//	beego.Router("/admin/category/", &controllers.CategoryAdminController{}, "*:Index")
//	beego.Router("/admin/category/search", &controllers.CategoryAdminController{}, "*:Search")
//	beego.Router("/admin/category/update", &controllers.CategoryAdminController{}, "*:Update")
//	beego.Router("/admin/category/view", &controllers.CategoryAdminController{}, "*:View")
}
