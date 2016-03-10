package controllers

import (
	models "app/models"
	"github.com/astaxie/beego"
)

type AdminController struct {
	BaseController
}

// 首页显示函数
func (this *AdminController) Get() {
	this.TplNames = "Layout/index.html"
}

// @router /Index [get]
func (this *AdminController) Index() {
	this.TplNames = "Admin/admin.html"
}

// ajax返回数据
func (this *AdminController) AjaxIndex() {
	// 默认返回
	Point := this.InitPoint()

	// 定义查询map类型
	mMap := map[string]string{
		"Status": "status",
		"Email":  "email__contains",
		"search": "username__contains",
		"Id":     "id",
	}

	// 查询字符串
	tmpMap, offset, limit, order := this.GetQueryString(mMap, "id")

	if this.IsAjax() {
		Point.Message = "获取数据为空"
		total, count, data, err := models.AdminGetAll(tmpMap, offset, limit, order)
		if err == nil {
			Point.Status = 1
			Point.Message = "success"
			Point.Data = this.DataTable(total, count, data)
		}
	}

	this.AjaxReturn(Point)
}

// 响应其他操作
func (this *AdminController) Save() {
	// 定义错误
	Point := this.InitPoint()

	// 接收参数
	actionType, admin, isHave := this.GetString("actionType"), models.Admin{}, true
	var Id, RowCount int64

	// 判断提交数据
	if actionType != "" {

		Point.Message = "数据赋值失败..."
		err := this.ParseForm(&admin)
		if err == nil {
			Point.Message = "添加管理员密码不能为空"
			if admin.Password != "" || actionType != "insert" {
				Point.Message = "账号密码和确认密码错误"
				Point.Data = admin

				// 验证确认密码和密码操作
				if admin.Password == this.GetString("Truepass") || actionType == "delete" {
					// 判断类型操作
					switch actionType {
					case "delete":
						isHave, RowCount, err = models.AdminDelete(admin.Id)
					case "insert":
						admin.LastIp = this.Ctx.Request.RemoteAddr
						Id, err = models.AdminInsert(&admin)
					case "update":
						isHave, RowCount, err = models.AdminSave(&admin)
					}

					// 判断操作是否成功
					Point.Message = "抱歉!没有查询到数据..."
					if isHave {
						Point.Message = "服务器繁忙,请稍候再试..."
						if err == nil || Id > 0 || RowCount > 0 {
							Point.Status = 1
							Point.Message = "恭喜你！操作成功 ^.^ "
						}
					}
				}
			}

		} else {
			beego.Error(err)
		}
	}

	this.AjaxReturn(Point)
}
