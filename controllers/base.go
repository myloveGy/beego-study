package controllers

import (
	models "app/models"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	// "time"
)

// 定义公共的控制器
type BaseController struct {
	beego.Controller
}

// 控制器的前置操作
func (this *BaseController) Prepare() {
	admin := this.GetSession("AdminUser")
	// // 验证用户有没有登录
	// if admin == nil {
	// 	this.Redirect("/", 302)
	// }

	// 默认注册
	if admin != nil {
		user := admin.(models.Admin)
		this.Data["admin"] = user
	}
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

// 获取查询数据信息
func (this *BaseController) GetQueryString() (tmpMap map[string]string) {
	str := this.GetString("msearch")
	if str != "" {
		json, _ := simplejson.NewJson([]byte(str))
		array, _ := json.Array()
		tmpMap := map[string]string{}
		for key, _ := range array {
			name, _ := json.GetIndex(key).Get("name").String()
			value, _ := json.GetIndex(key).Get("value").String()
			tmpMap[name] = value
		}
	}

	return

}
