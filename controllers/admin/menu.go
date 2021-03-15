package admin

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego/orm"

	"project/models"
)

type Menu struct {
	Comm
}

// 首页显示
func (m *Menu) Index() {
	// 查询主要导航
	var arr []*models.Menu
	if _, err := orm.NewOrm().QueryTable(&models.Menu{}).All(&arr); err == nil {
		data := make(map[string]string)
		data["0"] = "顶级分类"
		for _, v := range arr {
			data[strconv.FormatInt(v.Id, 10)] = v.MenuName
		}
		str, _ := json.Marshal(&data)
		m.Data["menus"] = string(str)
	}

	m.TplName = "admin/menu/index.html"
}

// 查询数据
func (m *Menu) Search() {
	arr := make([]*models.Menu, 0)

	// 查询信息
	search := map[string]string{
		"menu_name": "menu_name__icontains",
		"id":        "id",
		"status":    "status",
		"url":       "url__icontains",
	}

	// 返回信息
	m.baseSearch(&arr, search)
}

// 修改数据
func (m *Menu) Update() {
	m.baseUpdate(&models.Menu{})
}

// Delete 删除数据
func (m *Menu) Delete() {
	m.baseDelete(&models.Menu{})
}
