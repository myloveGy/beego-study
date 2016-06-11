package controllers

import (
	"encoding/json"
	"project/models"
	"strconv"

	//	"github.com/astaxie/beego"
)

type CategoryAdminController struct {
	AdminController
}

// 首页显示
func (this *CategoryAdminController) Index() {
	// 查询分类的顶级分类
	query := models.QueryOther{
		Table: "my_category",
		Where: map[string]interface{}{
			"status": 1,
			"pid":    0,
		},
	}
	var arr []*models.Category
	if _, err := models.All(&arr, query); err == nil {

		data := make(map[string]string)
		data["0"] = "顶级分类"
		for _, v := range arr {
			data[strconv.FormatInt(v.Id, 10)] = v.CateName
		}
		str, _ := json.Marshal(&data)
		this.Data["categorys"] = string(str)
	}
	this.TplName = "admin/category.html"
}

// 查询数据
func (this *CategoryAdminController) Search() {
	var arr []*models.Category

	// 查询信息
	search := map[string]string{
		"search":  "cate_name__icontains",
		"id":      "id",
		"status":  "status",
		"orderBy": "id",
		"Table":   "my_category",
	}

	// 返回信息
	this.BaseSearch(&arr, search)
}

// 修改数据
func (this *CategoryAdminController) Update() {
	var object models.Category
	this.BaseUpdate(&object, "my_category")
}
