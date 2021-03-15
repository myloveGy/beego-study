package admin

import (
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"

	"project/controllers"
	"project/logic"
	"project/repositories"
	"project/response"
	"project/utils"
)

// 后台控制器
type Comm struct {
	controllers.Base
}

// 前置操作
func (c *Comm) Prepare() {

	// 如果是ajax 请求获取POST提交、忽略
	if c.IsAjax() || c.Ctx.Request.Method == "POST" {
		if !c.IsLogin("admin") {
			response.NotLogin(&c.Base.Controller, "还没有登录")
			return
		}

		return
	}

	// 没有登录
	if !c.IsLogin("admin") {
		c.Redirect("/admin", 302)
		return
	}

	// 使用的布局
	c.Data["admin"] = c.User
	c.Data["navigation"] = logic.GetCacheMenu()
	c.Layout = "admin/layout/main.html"
}

// 查询方法
func (c *Comm) Query(search map[string]string) repositories.Query {
	query := new(repositories.Query)

	// 处理默认查询信息
	query.Table = search["Table"]
	query.Order = search["orderBy"]
	query.IStart, _ = c.GetInt64("iDisplayStart")
	query.ILength, _ = c.GetInt64("iDisplayLength")
	query.Where = make(map[string]interface{})

	// 判断排序字段
	if order := c.GetString("params[orderBy]"); order != "" {
		query.Order = order
		delete(c.Ctx.Request.PostForm, "params[orderBy]")
	}

	// 判断排序方式
	sType := c.GetString("sSortDir_0")
	if sType != "" {
		query.Order = strings.TrimLeft(query.Order, "-")
		if sType == "desc" {
			query.Order = "-" + query.Order
		}

		delete(c.Ctx.Request.PostForm, "sSortDir_0")
	}

	// 判断查询信息
	if request := c.Ctx.Request.PostForm; request != nil {
		// 取出其他查询条件
		for k, v := range request {
			if strings.HasPrefix(k, "params[") {
				key := strings.Trim(strings.Trim(strings.Trim(k, "params"), "]"), "[")
				if tmp, ok := search[key]; ok {
					query.Where[tmp] = v[0]
				}
			}
		}
	}

	return *query
}

// 公共的查询数据的方法
func (c *Comm) BaseSearch(arr interface{}, search map[string]string, where map[string]interface{}) {
	// 定义返回数据
	var (
		data response.DataTable
		err  error
	)

	// 处理查询数据信息
	query := c.Query(search)
	for k, v := range where {
		query.Where[k] = v
	}

	// 查询数据
	data.RecordsTotal, _, err = repositories.FindAll(arr, query)
	if err != nil {
		response.BusinessError(&c.Base.Controller, "服务器繁忙，请稍后再试")
		return
	}

	data.Data = arr
	response.Success(&c.Base.Controller, &data)
}

// 公共的编辑的方法
func (c *Comm) BaseUpdate(object interface{}, table string) {
	// 获取请求信息
	actionType := c.GetString("actionType")
	if actionType == "" {
		response.MissingParams(&c.Base.Controller)
		return
	}

	if !utils.InArray([]string{"insert", "update", "delete", "deleteAll"}, actionType) {
		response.MissingParams(&c.Base.Controller, "请求类型错误")
		return
	}

	// 修改数据
	if actionType == "update" {
		// 修改数据需要先查询数据
		id, err := c.GetInt64("id")
		if err != nil {
			response.BusinessError(&c.Base.Controller, "主键数据不存在")
			return
		}

		if err := repositories.One(object, repositories.QueryOther{Table: table, Where: map[string]interface{}{"id": id}}); err != nil {
			response.BusinessError(&c.Base.Controller, "修改数据不存在")
			return
		}
	}

	if actionType == "deleteAll" {
		c.baseDeleteAll(object, table)
		return
	}

	if err := c.ParseForm(object); err != nil {
		response.InvalidParams(&c.Base.Controller, "请求数据为空")
		return
	}

	var err error
	// 根据类型做出相应的处理
	switch actionType {
	case "insert": // 新增数据
		_, err = orm.NewOrm().Insert(object)
	case "update": // 修改数据
		_, err = orm.NewOrm().Update(object)
	case "delete": // 删除数据
		_, err = repositories.Delete(object)
	}

	// 判断返回数据
	if err != nil {
		response.BusinessError(&c.Base.Controller, "抱歉！执行该操作出现错误")
		return
	}

	response.Success(&c.Base.Controller, object, "操作成功")
}

func (c *Comm) baseDelete(data interface{}, table string) {
	// 获取主键
	strId := "id"
	if v, ok := data.(repositories.Model); ok {
		strId = v.PK()
	}

	// 修改数据需要先查询数据
	id, err := c.GetInt64(strId)
	if err != nil {
		response.MissingParams(&c.Base.Controller, "主键数据不存在")
		return
	}

	if err := repositories.One(data, repositories.QueryOther{Table: table, Where: map[string]interface{}{strId: id}}); err != nil {
		response.BusinessError(&c.Base.Controller, "删除数据不存在")
		return
	}

	if _, err := repositories.Delete(data); err != nil {
		response.BusinessError(&c.Base.Controller, "抱歉！删除数据出现错误")
		return
	}

	response.Success(&c.Base.Controller, data, "操作成功")
}

// BaseDeleteAll 批量删除
func (c *Comm) baseDeleteAll(data interface{}, table string) {
	ids := c.GetString("ids")
	if ids == "" {
		response.MissingParams(&c.Base.Controller, "删除数据为空")
		return
	}

	aIds := strings.Split(ids, ",")
	if len(aIds) == 0 {
		response.BusinessError(&c.Base.Controller, "删除数据为空")
		return
	}

	if _, err := repositories.DeleteAll(data, aIds, table); err != nil {
		response.SystemError(&c.Base.Controller, "删除数据失败")
		return
	}

	response.Success(&c.Base.Controller, aIds, "批量删除成功")
}

// BaseUpload 图片上传处理
func (c *Comm) BaseUpload(filename, pathname string, allowType []string, size int64, oldFile string) {
	oldPath := c.GetString(oldFile)
	if oldPath != "" {
		// 删除之前的文件
		_ = os.Remove("." + oldPath)
	}

	f, h, err := c.GetFile(filename)
	if err != nil {
		response.MissingParams(&c.Base.Controller, "没有文件上传")
		return
	}

	defer f.Close()
	file := path.Ext(h.Filename)
	if !utils.InArray(allowType, file) {
		response.InvalidParams(&c.Base.Controller, "上传文件格式不对")
		return
	}

	if size < f.(utils.Sizer).Size() {
		response.InvalidParams(&c.Base.Controller, "上传文件不能超过过大")
		return
	}

	// 处理上传目录
	dirName := "./static/uploads/" + pathname
	// 目录不存在创建
	if !utils.IsDirExists(dirName) {
		if err = os.MkdirAll(dirName, 0777); err != nil {
			response.BusinessError(&c.Base.Controller, "创建目录失败 :( "+dirName)
			return
		}
	}

	// 文件最终保存的地址
	fileName := dirName + "/" + strconv.Itoa(rand.Int()) + file
	if err = c.SaveToFile(filename, fileName); err != nil {
		response.BusinessError(&c.Base.Controller, "上传失败")
		return
	}

	data := map[string]string{"fileName": h.Filename, "fileUrl": fileName[1:]}
	response.Success(&c.Base.Controller, data, "文件上传成功")
}
