package controllers

import (
	models "app/models"
)

type MenusController struct {
	BaseController
}

func (this *MenusController) Index() {
	this.TplNames = "Layout/base.html"
}

// 响应ajax获取数据
func (this *MenusController) AjaxIndex() {
	// 默认返回
	status := 0
	sEcho, _ := this.GetInt64("sEcho", 1)
	source := make(map[string]interface{})
	var message string
	message = "提交参数问题"
	if this.IsAjax() {
		message = "获取数据为空"
		var err error
		source["iTotalDisplayRecords"], source["aaData"], err = models.MenusGetAll()
		source["iTotalRecords"] = 0
		source["sEcho"] = sEcho
		if err == nil {
			status = 1
		}

	}

	this.AjaxReturn(status, message, source)
}

// 响应其他操作
func (this *MenusController) Update() {
	this.AjaxReturn(0, "提交参数错误", nil)
}
