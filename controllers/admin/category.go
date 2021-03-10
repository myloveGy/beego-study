package admin

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"project/controllers"
	"project/models"
)

type CategoryController struct {
	CommController
}

// 首页显示
func (c *CategoryController) Index() {
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
		c.Data["categorys"] = string(str)
	}
	c.TplName = "admin/category.html"
}

// 查询数据
func (c *CategoryController) Search() {
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
	c.BaseSearch(&arr, search, map[string]interface{}{"pid": 0})
}

// 修改数据
func (c *CategoryController) Update() {
	c.BaseUpdate(&models.Category{}, "my_category")
}

// 详情信息
func (c *CategoryController) View() {
	// 获取ID
	id, err := c.GetInt64("id")
	if err != nil {
		c.Error(controllers.CodeMissingParams, "请求数据为空", nil)
		return
	}

	num, err1 := c.GetInt("sEcho", 0)
	if err1 != nil {
		c.Error(controllers.CodeMissingParams, "请求数据为空", nil)
		return
	}

	var (
		arr   []*models.Category
		total int64
	)

	query := models.QueryOther{Table: "my_category", Where: map[string]interface{}{"pid": id, "status": 1}}
	logs.Alert(query)
	total, err = models.All(&arr, query)
	if err != nil {
		c.Error(controllers.CodeBusinessError, "查询出现问题", nil)
		return
	}

	c.Success(&DataTable{Total: total, Count: total, Echo: num, Data: arr}, "")
}

// Inline 行内编辑执行
func (c *CategoryController) Inline() {

	// 获取ID
	name, value := c.GetString("name"), c.GetString("value")
	if name == "" || value == "" {
		c.Error(controllers.CodeMissingParams, "请求数据为空", nil)
		return
	}

	id, err := c.GetInt64("pk")
	if err != nil {
		c.Error(controllers.CodeMissingParams, "主键信息为空", nil)
		return
	}

	var cate models.Category
	if err := models.One(&cate, models.QueryOther{Table: "my_category", Where: map[string]interface{}{"id": id}}); err != nil {
		c.Error(controllers.CodeInvalidParams, "修改数据为空", nil)
		return
	}

	v := reflect.ValueOf(&cate)
	// 首字母大写
	name = strings.ToUpper(name[0:1]) + name[1:]
	tempName := v.Elem().FieldByName(name)
	if !tempName.IsValid() {
		c.Error(controllers.CodeInvalidParams, "修改字段不存在", nil)
		return
	}

	var tv interface{}
	switch tempName.Interface().(type) {
	case int:
		tv, err = strconv.Atoi(value)
		logs.Alert(tv, " type is int")
	case int64:
		tv, err = strconv.ParseInt(value, 10, 64)
		logs.Alert(tv, " type is int64")
	case string:
		tv = value
		logs.Alert(tv, " type is string")
	default:
		err = errors.New("数据类型不确定")
	}

	if err != nil {
		c.Error(controllers.CodeBusinessError, "服务器处理出现错误: "+err.Error(), nil)
		return
	}

	if tempName.CanSet() {
		tempName.Set(reflect.ValueOf(tv))
		if _, err = orm.NewOrm().Update(&cate, name); err != nil {
			c.Error(controllers.CodeBusinessError, "修改失败", nil)
			return
		}
	}

	c.Success(&cate, "修改成功")
}
