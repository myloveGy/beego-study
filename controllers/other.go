package controllers

type OtherController struct {
	BaseController
}

// 首页显示函数
func (this *OtherController) Index() {
	this.TplNames = "Other/index.html"
}
