package controllers

import (
	models "app/models"
	"github.com/astaxie/beego"
)

type MenusController struct {
	BaseController
}

func (this *MenusController) Index() {
	this.TplNames = "Layout/base.html"
}

// 响应ajax获取数据
func (this *MenusController) AjaxIndex() {
	source := make(map[string]interface{})
	// 默认返回
	status := 0
	sEcho, _ := this.GetInt64("sEcho", 1)

	// 查询字符串
	query := this.GetQueryString()

	// 开始位置和结束位置
	start, _ := this.GetInt64("iDisplayStart", 0)
	length, _ := this.GetInt64("iDisplayLength", 10)

	var message string
	message = "提交参数问题"
	if this.IsAjax() {
		message = "获取数据为空"
		var err error
		source["iTotalDisplayRecords"], source["aaData"], err = models.MenusGetAll(query, start, length)
		source["iTotalRecords"] = 0
		source["sEcho"] = sEcho
		source["query"] = query
		if err == nil {
			status = 1
		}

	}

	this.AjaxReturn(status, message, source)
}

// 响应其他操作
func (this *MenusController) Save() {
	// 接收参数
	actionType, message, status, IsHave, menus := this.GetString("actionType"), "提交参数错误", 0, false, models.Menus{}
	var Id, RowCount int64

	// 判断提交数据
	if actionType != "" {
		message = "数据赋值失败..."
		if err := this.ParseForm(&menus); err == nil {
			// 获取SESSION值
			admin := this.GetSession("AdminUser")

			// 获取修改用户ID
			if admin != nil {
				user := admin.(models.Admin)
				menus.UpdateId = user.Id
			}

			// 判断类型操作
			switch actionType {
			case "delete":
				IsHave, RowCount, err = models.MenusDelete(menus.Id)
			case "insert":
				IsHave = true
				Id, err = models.MenusInsert(&menus)
				if err != nil {
					beego.Error(err)
				}
			case "update":
				IsHave, RowCount, err = models.MenusSave(&menus)
			}

			// 判断操作是否成功
			message = "抱歉!没有查询到数据..."
			if IsHave {
				message = "抱歉！服务器繁忙,请稍候再试..."
				if RowCount > 0 || Id > 0 {
					status = 1
					message = "恭喜你！操作成功 ^.^ "
				}
			}
		}
	}

	this.AjaxReturn(status, message, menus)
}
