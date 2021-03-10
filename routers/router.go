package routers

import (
	"project/controllers"
	"project/controllers/admin"
	"project/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	// 前台显示
	beego.Router("/", &controllers.IndexController{})
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.ArticleController{})

	// 使用命名空间
	userNamespace := beego.NewNamespace("/user",
		beego.NSAutoRouter(&user.ImageController{}),
		beego.NSAutoRouter(&user.ArticleController{}),
	)

	// 后台显示
	beego.Router("/admin", &admin.SiteController{}, "*:Index")

	// 使用命名空间
	adminNamespace := beego.NewNamespace("/admin",
		beego.NSAutoRouter(&admin.SiteController{}),
		beego.NSAutoRouter(&admin.Controller{}),
		beego.NSAutoRouter(&admin.MenuController{}),
		beego.NSAutoRouter(&admin.CategoryController{}),
	)

	beego.AddNamespace(adminNamespace, userNamespace)
}
