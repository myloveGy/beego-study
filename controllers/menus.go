package controllers

import (
	models "app/models"
	// "fmt"
	"github.com/astaxie/beego"
)

type MenusController struct {
	BaseController
}

func (this *MenusController) Index() {
	this.TplNames = "Admin/menus.html"
}

// 响应ajax获取数据
func (this *MenusController) AjaxIndex() {
	// 默认返回
	Point := this.InitPoint()

	// 定义查询map类型
	mMap := map[string]string{
		"Url":    "url__contains",
		"Status": "status",
		"search": "menuname__contains",
		"Id":     "id",
	}

	// 查询字符串
	tmpMap, offset, limit, order := this.GetQueryString(mMap, "id")

	if this.IsAjax() {
		Point.Message = "获取数据为空"
		total, count, data, err := models.MenusGetAll(tmpMap, offset, limit, order)
		if err == nil {
			Point.Status = 1
			Point.Message = "success"
			Point.Data = this.DataTable(total, count, data)
		}

	}

	this.AjaxReturn(Point)
}

// 响应其他操作
func (this *MenusController) Save() {
	// 定义错误
	Point := this.InitPoint()

	// 接收参数
	actionType, IsHave, menus := this.GetString("actionType"), false, models.Menus{}
	var Id, RowCount int64

	// 判断提交数据
	if actionType != "" {
		Point.Message = "数据赋值失败..."
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
			Point.Message = "抱歉!没有查询到数据..."
			if IsHave {
				Point.Message = "抱歉！服务器繁忙,请稍候再试..."
				if RowCount > 0 || Id > 0 {
					Point.Status = 1
					Point.Message = "恭喜你！操作成功 ^.^ "
				}
			}
		}
	}

	this.AjaxReturn(Point)
}
