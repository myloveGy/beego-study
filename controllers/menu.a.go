package controllers

import (
	"project/models"
)

type MenuAdminController struct {
	AdminController
}

// 首页显示
func (this *MenuAdminController) Index() {
	this.TplName = "admin/menu.html"
}

func (this *MenuAdminController) Search() {
	this.E = ArrError{Status: 1, Msg: "Success", Data: nil}
	arr, _ := models.SearchMenu(map[string]interface{}{}, 10, 0, "id")
	this.E.Data = DataTable{
		Echo:  0,
		Count: len(arr),
		Total: int64(len(arr)),
		Data:  arr,
	}
	this.AjaxReturn()
}
