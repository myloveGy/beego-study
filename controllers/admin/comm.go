package admin

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/jinxing-go/mysql"

	"project/connection"
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

func (c *Comm) baseCreate(data mysql.Model) {
	// 解析出数据
	if err := c.ParseForm(data); err != nil {
		response.InvalidParams(&c.Base.Controller, "请求数据为空")
		return
	}

	// 执行前置操作
	if v, ok := data.(models.BeforeSave); ok {
		if err := v.BeforeSave(); err != nil {
			response.BusinessError(&c.Base.Controller, err)
			return
		}
	}

	// 新增数据
	if err := connection.DB.Create(data); err != nil {
		response.BusinessError(&c.Base.Controller, err)
		return
	}

	response.Success(&c.Base.Controller, data)
}

// 公共的查询数据的方法
func (c *Comm) baseSearch(data interface{}, search map[string]string) {

	// 接收参数
	draw, _ := c.GetInt64("draw", 1)
	limit, _ := c.GetInt64("limit", 10)
	offset, _ := c.GetInt64("offset", 0)
	page := offset / limit

	// 处理排序
	orderBy := c.GetString("orderBy")
	order := strings.Split(orderBy, " ")

	// 查询数据
	resp := &response.DataTable{Draw: draw}
	query := connection.DB.Builder(data)

	// 添加查询条件
	for k, v := range search {
		if k == "default" {
			where := strings.Split(v, ",")
			for i := 0; i < len(where); i += 2 {
				query = query.Where(where[i], where[i+1])
			}
		} else {
			value := c.Ctx.Input.Query(fmt.Sprintf("filters[%s]", k))
			if value != "" && v != "" {
				if v == "like" {
					value = "%" + value + "%"
				}

				query = query.Where(k, v, value)
			}
		}
	}

	// 查询数据总条数
	total, err := query.OrderBy(order[0], order[1]).Paginate(int(page), int(limit))
	if err != nil {
		response.BusinessError(&c.Base.Controller, "服务器繁忙，请稍后再试")
		return
	}

	resp.Data = data
	resp.RecordsTotal = total
	resp.RecordsFiltered = total
	response.Success(&c.Base.Controller, resp)
}

// 公共的编辑的方法
func (c *Comm) baseUpdate(object mysql.Model) {

	// 查询数据是否存在
	if err := c.findOrFail(object, false); err != nil {
		response.InvalidParams(&c.Base.Controller, "抱歉！修改数据不存在")
		return
	}

	// 解析数据信息
	if err := c.ParseForm(object); err != nil {
		response.InvalidParams(&c.Base.Controller, "请求数据为空")
		return
	}

	// 执行前置操作
	if v, ok := object.(models.BeforeSave); ok {
		if err := v.BeforeSave(); err != nil {
			response.BusinessError(&c.Base.Controller, err)
			return
		}
	}

	// 执行修改数据
	if _, err := connection.DB.Update(object); err != nil {
		response.BusinessError(&c.Base.Controller, "抱歉！修改数据出现错误")
		return
	}

	response.Success(&c.Base.Controller, object, "操作成功")
}

func (c *Comm) baseDelete(data mysql.Model) {

	// 查询数据是否存在
	if err := c.findOrFail(data, true); err != nil {
		response.InvalidParams(&c.Base.Controller, "抱歉！删除数据不存在")
		return
	}

	// 执行删除数据
	if _, err := connection.DB.Delete(data); err != nil {
		response.BusinessError(&c.Base.Controller, "抱歉！删除数据出现错误")
		return
	}

	response.Success(&c.Base.Controller, data, "操作成功")
}

func (c *Comm) findOrFail(data mysql.Model, isFind bool) error {
	// 获取主键
	strId := data.PK()

	// 修改数据需要先查询数据
	id, err := c.GetInt64(strId)
	if err != nil {
		return err
	}

	if isFind {
		if err := connection.DB.Builder(data).Where(strId, id).One(); err != nil {
			return errors.New("数据不存在")
		}

		return nil
	}

	var total int64
	if err := connection.DB.Get(&total, fmt.Sprintf("SELECT COUNT(*) AS `total` FROM `%s` WHERE `%s` = ? ", data.TableName(), strId), id); err != nil {
		return errors.New("数据不存在")
	}

	if total == 0 {
		return errors.New("数据不存在")
	}

	return nil
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
