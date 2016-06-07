package controllers

type MenuAdminController struct {
	AdminController
}

func (this *MenuAdminController) Index() {
	this.AjaxReturn()
}
