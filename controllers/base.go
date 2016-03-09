package controllers

import (
	models "app/models"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	// "strings"
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

/**
 * GetQueryString() 处理查询信息
 * @param map 查询字段对应关系
 * @return map int64 in64
 */
func (this *BaseController) GetQueryString(mMap map[string]string, sortBy string) (tmpMap map[string]interface{}, offset int64, limit int64, orderBy string) {
	// 查询的字段信息
	str := this.GetString("msearch")
	if str != "" {
		json, _ := simplejson.NewJson([]byte(str))
		array, _ := json.Array()
		tmpMap = map[string]interface{}{}
		for key, _ := range array {
			name, _ := json.GetIndex(key).Get("name").String()
			if mMap[name] != "" {
				value, err := json.GetIndex(key).Get("value").String()
				temp := mMap[name]
				name = temp
				tmpMap[name] = value
				if err != nil {
					beego.Error(err)
				}
			}

			// 排序字段
			if name == "orderBy" {
				sortBy, _ = json.GetIndex(key).Get("value").String()
			}
		}
	}

	// 排序方式
	if this.GetString("sSortDir_0", "asc") == "desc" {
		orderBy = "-" + sortBy
	} else {
		orderBy = sortBy
	}

	// 开始位置和结束位置
	offset, _ = this.GetInt64("iDisplayStart", 0)
	limit, _ = this.GetInt64("iDisplayLength", 10)

	return
}
