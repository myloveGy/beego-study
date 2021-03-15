package admin

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"

	"project/controllers"
	"project/logic"
	"project/models"
	"project/response"
	"project/utils"
)

// 后台控制器
type Comm struct {
	controllers.Base
}

// 前置操作
func (c *Comm) Prepare() {

	// // 如果是ajax 请求获取POST提交、忽略
	// if c.IsAjax() || c.Ctx.Request.Method == "POST" {
	// 	if !c.IsLogin("admin") {
	// 		response.NotLogin(&c.Base.Controller, "还没有登录")
	// 		return
	// 	}
	//
	// 	return
	// }
	//
	// // 没有登录
	// if !c.IsLogin("admin") {
	// 	c.Redirect("/admin", 302)
	// 	return
	// }

	// 使用的布局
	c.Data["admin"] = c.User
	c.Data["navigation"] = logic.GetCacheMenu()
	c.Layout = "admin/layout/main.html"
}

func (c *Category) baseCreate(data interface{}) {
	if err := c.ParseForm(data); err != nil {
		response.InvalidParams(&c.Base.Controller, "请求数据为空")
		return
	}

	if _, err := orm.NewOrm().Insert(data); err != nil {
		response.BusinessError(&c.Base.Controller, err)
		return
	}

	response.Success(&c.Base.Controller, data)
}

// 公共的查询数据的方法
func (c *Comm) baseSearch(data interface{}, search map[string]string) {

	// 获取model信息
	modelObject, err := models.GetModel(data)
	if err != nil {
		response.BusinessError(&c.Base.Controller, "服务器繁忙，请稍后再试")
		return
	}

	// 接收参数
	draw, _ := c.GetInt64("draw", 1)
	limit, _ := c.GetInt64("limit", 10)
	offset, _ := c.GetInt64("offset", 0)

	// 处理排序
	orderBy := c.GetString("orderBy", fmt.Sprintf("%s desc", modelObject.PK()))
	order := strings.Split(orderBy, " ")
	if len(order) == 2 && order[1] == "acs" {
		orderBy = fmt.Sprintf("-%s", order[0])
	} else {
		orderBy = fmt.Sprintf("%s", order[0])
	}

	// 查询数据
	resp := &response.DataTable{Draw: draw}
	table := modelObject.TableName()
	query := orm.NewOrm().QueryTable(table)

	// 添加查询条件
	for k, v := range search {
		if k == "default" {
			where := strings.Split(v, ",")
			for i := 0; i < len(where); i += 2 {
				query = query.Filter(where[i], where[i+1])
			}
		} else {
			value := c.Ctx.Input.Query(fmt.Sprintf("filters[%s]", k))
			if value != "" && v != "" {
				query = query.Filter(v, value)
			}
		}
	}

	// 查询数据总条数
	if resp.RecordsTotal, err = query.Count(); err != nil {
		response.BusinessError(&c.Base.Controller, "服务器繁忙，请稍后再试")
		return
	}

	// 查询数据信息
	if _, err := query.Offset(offset).OrderBy(orderBy).Limit(limit).All(data); err != nil {
		response.BusinessError(&c.Base.Controller, "服务器繁忙，请稍后再试")
		return
	}

	resp.Data = data
	resp.RecordsFiltered = resp.RecordsTotal
	response.Success(&c.Base.Controller, resp)
}

// 公共的编辑的方法
func (c *Comm) baseUpdate(object interface{}) {

	// 查询数据是否存在
	if err := c.findOrFail(object); err != nil {
		response.InvalidParams(&c.Base.Controller, "抱歉！修改数据不存在")
		return
	}

	// 解析数据信息
	if err := c.ParseForm(object); err != nil {
		response.InvalidParams(&c.Base.Controller, "请求数据为空")
		return
	}

	// 执行修改数据
	if _, err := orm.NewOrm().Update(object); err != nil {
		response.BusinessError(&c.Base.Controller, "抱歉！修改数据出现错误")
		return
	}

	response.Success(&c.Base.Controller, object, "操作成功")
}

func (c *Comm) baseDelete(data interface{}) {

	// 查询数据是否存在
	if err := c.findOrFail(data); err != nil {
		response.InvalidParams(&c.Base.Controller, "抱歉！删除数据不存在")
		return
	}

	// 执行删除数据
	if _, err := orm.NewOrm().Delete(data); err != nil {
		response.BusinessError(&c.Base.Controller, "抱歉！删除数据出现错误")
		return
	}

	response.Success(&c.Base.Controller, data, "操作成功")
}

func (c *Comm) findOrFail(data interface{}) error {
	// 获取主键
	strId := "id"
	if v, ok := data.(models.Model); ok {
		strId = v.PK()
	}

	// 修改数据需要先查询数据
	id, err := c.GetInt64(strId)
	if err != nil {
		return err
	}

	return orm.NewOrm().QueryTable(data).Filter(strId, id).One(data)
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
