package controllers

import (
	models "app/models"
	"github.com/astaxie/beego"
	// "time"
)

// 定义公共的控制器
type BaseController struct {
	beego.Controller
}

// 控制器的前置操作
func (this *BaseController) Prepare() {
	// admin := this.GetSession("AdminUser")
	// // 验证用户有没有登录
	// if admin == nil {
	// 	this.Redirect("/", 302)
	// }

	// // 没有注册信息到模板中
	// user := admin.(models.Admin)
	// this.Data["admin"] = user
	// this.Data["CreateTime"] = time.Unix(user.CreateTime, 0).Format("2006-01-02 03:04:05 PM")
	// this.Data["LastTime"] = time.Unix(user.LastTime, 0).Format("2006-01-02 03:04:05 PM")
	admin := models.Admin{Username: "admin"}
	this.Data["admin"] = admin
}

// ajax返回数据
func (this *BaseController) AjaxReturn(status int, message string, data interface{}) {
	returnData := struct {
		Status  int
		Message string
		Data    interface{}
	}{
		Status:  status,
		Message: message,
		Data:    data,
	}

	this.Data["json"] = returnData
	this.ServeJson()
}
