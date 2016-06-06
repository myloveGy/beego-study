package controllers

// 首页控制器
type AdminController struct {
	BaseController
}

func (this *AdminController) Prepare() {
	// 判断用户是否已经登录, 没有登录返回到登录页面
	if !this.isLogin("admin") {
		this.Redirect("/admin", 302)
	}

	// 使用的布局
	this.Data["admin"] = this.U
	this.Layout = "layout/admin.tpl"
}
