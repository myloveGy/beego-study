package admin

import (
	"encoding/json"
	"strconv"

	"project/models"
)

type MenuController struct {
	CommController
}

// 首页显示
func (m *MenuController) Index() {
	// 查询主要导航
	query := models.QueryOther{
		Table: "my_menu",
		Where: map[string]interface{}{
			"status": 1,
			"pid":    0,
		},
	}
	var arr []*models.Menu
	if _, err := models.All(&arr, query); err == nil {
		data := make(map[string]string)
		data["0"] = "顶级分类"
		for _, v := range arr {
			data[strconv.FormatInt(v.Id, 10)] = v.MenuName
		}
		str, _ := json.Marshal(&data)
		m.Data["menus"] = string(str)
	}

	m.TplName = "admin/menu.html"
}

// 查询数据
func (m *MenuController) Search() {
	var arr []*models.Menu

	// 查询信息
	search := map[string]string{
		"search":  "menu_name__icontains",
		"id":      "id",
		"status":  "status",
		"url":     "url__icontains",
		"orderBy": "id",
		"Table":   "my_menu",
	}

	// 返回信息
	m.BaseSearch(&arr, search, map[string]interface{}{})
}

// 修改数据
func (m *MenuController) Update() {
	m.BaseUpdate(&models.Menu{}, "my_menu")
}
