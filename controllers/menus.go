package controllers

type MenusController struct {
	BaseController
}

func (this *MenusController) Index() {
	this.TplNames = "Layout/base.html"
}
