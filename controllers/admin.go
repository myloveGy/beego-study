package controllers

// 表格返回数据
type DataTable struct {
	Echo  int         `json:"sEcho"`
	Count int         `json:"iTotalRecords"`
	Total int64       `json:"iTotalDisplayRecords"`
	Data  interface{} `json:"aaData"`
}

// 后台控制器
type AdminController struct {
	BaseController
}

// 前置操作
func (this *AdminController) Prepare() {
	// 判断用户是否已经登录, 没有登录返回到登录页面
	if !this.isLogin("admin") {
		this.Redirect("/admin", 302)
	}

	// 使用的布局
	this.Data["admin"] = this.U
	this.Layout = "layout/admin.tpl"
}
