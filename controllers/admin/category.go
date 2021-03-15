package admin

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"

	"project/connection"
	"project/models"
	"project/response"
)

type Category struct {
	Comm
}

// 首页显示
func (c *Category) Index() {

	arrList := make([]*models.Category, 0)
	if err := connection.DB.
		Builder(&arrList).
		Where("status", 1).
		Where("pid", 0).
		All(); err == nil {
		data := make(map[string]string)
		data["0"] = "顶级分类"
		for _, v := range arrList {
			data[strconv.FormatInt(v.Id, 10)] = v.CateName
		}
		str, _ := json.Marshal(&data)
		c.Data["categoryList"] = string(str)
	}

	c.TplName = "admin/category/index.html"
}

func (c *Category) Create() {
	c.baseCreate(&models.Category{})
}

// 查询数据
func (c *Category) Search() {
	arr := make([]*models.Category, 0)

	// 查询信息
	search := map[string]string{
		"id":     "=",
		"status": "=",
	}

	// 返回信息
	c.baseSearch(&arr, search)
}

// Update 修改数据
func (c *Category) Update() {
	c.baseUpdate(&models.Category{})
}

// Delete 删除数据
func (c *Category) Delete() {
	c.baseDelete(&models.Category{})
}

// Inline 行内编辑执行
func (c *Category) Inline() {

	// 获取ID
	name, value := c.GetString("name"), c.GetString("value")
	if name == "" || value == "" {
		response.MissingParams(&c.Base.Controller, "请求数据为空")
		return
	}

	id, err := c.GetInt64("pk")
	if err != nil || id <= 0 {
		response.MissingParams(&c.Base.Controller, "主键信息为空")
		return
	}

	category := &models.Category{Id: id}
	if err := connection.DB.Find(category); err != nil {
		response.InvalidParams(&c.Base.Controller, "修改数据为空")
		return
	}

	updateCategory := &models.Category{Id: id}
	v := reflect.ValueOf(updateCategory)

	// 首字母大写
	name = strings.ToUpper(name[0:1]) + name[1:]
	tempName := v.Elem().FieldByName(name)
	if !tempName.IsValid() {
		response.InvalidParams(&c.Base.Controller, "修改字段不存在")
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
		response.BusinessError(&c.Base.Controller, "服务器处理出现错误: "+err.Error(), nil)
		return
	}

	if tempName.CanSet() {
		tempName.Set(reflect.ValueOf(tv))
		if _, err = connection.DB.Update(updateCategory, name); err != nil {
			response.BusinessError(&c.Base.Controller, "修改失败")
			return
		}
	}

	response.Success(&c.Base.Controller, category, "修改成功")
}
