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

// 导航栏信息
type Menus struct {
	Id       int64
	MenuName string
	Icons    string
	Url      string
	Child    []interface{}
	Len      int
}

// 控制器的前置操作
func (this *BaseController) Prepare() {
	admin := this.GetSession("AdminUser")
	// 验证用户有没有登录
	if admin == nil {
		this.Redirect("/", 302)
	}

	// 默认注册
	if admin != nil {
		user := admin.(models.Admin)
		this.Data["admin"] = user
	}

	// 查询导航信息
	menusAll := map[int64]Menus{}
	menus, err := models.GetAllMenus()
	if err == nil {

		// 处理目录信息
		for _, value := range menus {
			pid, id := value.Pid, value.Id
			if pid == 0 {
				// 一级目录
				if _, ok := menusAll[id]; ok == true {
					m := menusAll[id]
					m.Id = value.Id
					m.MenuName = value.MenuName
					m.Icons = value.Icons
					m.Url = value.Url
					m.Len = len(m.Child)
					menusAll[id] = m
				} else {
					menusAll[id] = Menus{
						Id:       value.Id,
						MenuName: value.MenuName,
						Icons:    value.Icons,
						Url:      value.Url,
					}
				}

			} else {
				// 二级目录
				if _, ok := menusAll[pid]; ok == true {
					m := menusAll[pid]
					m.Child = append(m.Child, Menus{
						Id:       value.Id,
						MenuName: value.MenuName,
						Icons:    value.Icons,
						Url:      value.Url,
					})

					m.Len = len(m.Child)
					menusAll[pid] = m
				} else {
					m := Menus{
						Id:       value.Id,
						MenuName: value.MenuName,
						Icons:    value.Icons,
						Url:      value.Url,
					}

					menusAll[pid] = Menus{Child: []interface{}{m}}
				}
			}
		}
	}

	// 注入变量
	this.Data["Menus"] = menusAll
}

// 提示信息
type MePoint struct {
	Status  int
	Message string
	Data    interface{}
}

// 初始化话定义返回数据
func (this *BaseController) InitPoint() MePoint {
	return MePoint{
		Status:  0,
		Message: "数据为空",
		Data:    nil,
	}
}

// ajax返回数据
func (this *BaseController) AjaxReturn(Point MePoint) {
	this.Data["json"] = Point
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

// 返回数据给DataTable
func (this *BaseController) DataTable(total int64, count int, data interface{}) (tmpMap map[string]interface{}) {
	sEcho, _ := this.GetInt64("sEcho", 1)
	// 初始化定义
	tmpMap = map[string]interface{}{
		"sEcho":                sEcho,
		"iTotalDisplayRecords": total,
		"iTotalRecords":        count,
		"aaData":               data,
	}

	return
}
