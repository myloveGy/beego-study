package controllers

type AdminController struct {
	BaseController
}

// 首页显示函数
func (this *AdminController) Get() {
	this.TplNames = "Layout/index.html"
}

// @router /Index [get]
func (this *AdminController) Index() {
	this.TplNames = "Layout/base.html"
}
