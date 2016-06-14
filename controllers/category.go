package controllers

import (
	"encoding/json"
	"project/models"
	"strconv"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	CommController
}

// 首页显示
func (this *CategoryController) Index() {
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
func (this *CategoryController) Search() {
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
	this.BaseSearch(&arr, search, map[string]interface{}{"pid":0,})
}

// 修改数据
func (this *CategoryController) Update() {
	var object models.Category
	this.BaseUpdate(&object, "my_category")
}

// 详情信息
func (this *CategoryController) View() {
	this.E = ArrError{Status: 0, Msg: "请求数据为空", Data: nil}
	// 获取ID
	id, err := this.GetInt64("id")
	s := 0
	if num, err := this.GetInt("sEcho"); err == nil {
		s = num
	}

	if err == nil {
		var arr []*models.Category
		query := models.QueryOther{Table:"my_category", Where:map[string]interface{}{"pid":id, "status":1}}
		beego.Alert(query)
		if num, err := models.All(&arr, query); err == nil {
			this.E.Data = DataTable{Total:num, Count:num, Echo:s, Data:arr}
			this.E.Msg  = "success"
			this.E.Status = 1
		}
	} else {
		this.E.Msg = "服务器处理出现错误 Error ：" + err.Error()
	}

	this.AjaxReturn()
}

// 行内编辑执行
func (this *CategoryController) Inline() {
	this.E = ArrError{Status: 0, Msg: "请求数据为空", Data: nil}
	var cate models.Category
	// 获取ID
	id, err := this.GetInt64("pk")
	if err == nil {
		err = models.One(&cate, models.QueryOther{Table:"my_category", Where:map[string]interface{}{"id":id}})
		if err == nil {
			err = this.ParseForm(&cate)
			if err == nil {
				id, err = models.Update(&cate)
				this.E.Msg = "服务器繁忙,请稍候再试..."
				if id > 0 && err == nil {
					this.E.Msg = "修改成功"
					this.E.Status = 1
					this.E.Data = cate
				} else {
					this.E.Msg = "服务器处理出现错误 Error ：" + err.Error()
				}
			}
		}
	} else {
		this.E.Msg = "服务器处理出现错误 Error ：" + err.Error()
	}



	this.AjaxReturn()
}
