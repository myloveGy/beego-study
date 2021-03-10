package admin

import (
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"

	"project/controllers"
	"project/help"
	"project/models"
)

// 表格返回数据
type DataTable struct {
	Echo  int         `json:"sEcho"`
	Count int64       `json:"iTotalRecords"`
	Total int64       `json:"iTotalDisplayRecords"`
	Data  interface{} `json:"aaData"`
}

// 后台控制器
type CommController struct {
	controllers.BaseController
	SearchMap func() map[string]string
}

type StuMenu struct {
	MenuName string                 `json:"MenuName"`
	Icons    string                 `json:"Icons"`
	Url      string                 `json:"Url"`
	Child    map[string]models.Menu `json:"Child"`
	IsChild  bool                   `json:"IsChild"`
}

type MeMenus map[string]StuMenu

// 前置操作
func (c *CommController) Prepare() {

	bTrue := true
	menus := make(MeMenus)
	oCache, err := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`)

	// 判断缓存对象存在并且有缓存数据信息
	if err == nil {
		if m := oCache.Get("menus"); m != nil && oCache.IsExist("menus") {
			switch m.(type) {
			case MeMenus:
				menus = m.(MeMenus)
				bTrue = false
			default:
				oCache.Delete("menus")
			}
		}
	}

	// 没有缓存数据信息 - 查询导航栏信息
	if bTrue {
		var menu []*models.Menu
		models.All(&menu, models.QueryOther{Table: "my_menu", Where: map[string]interface{}{"status": 1}, Order: "-sort"})
		for _, value := range menu {
			tmpKey := strconv.FormatInt(value.Id, 10)
			if value.Pid != 0 {
				tmpKey = strconv.FormatInt(value.Pid, 10)
			}
			// 判断数据是首页导航数据
			if value.Pid == 0 {

				if _, v := menus[tmpKey]; !v {
					ma := make(map[string]models.Menu)
					menus[tmpKey] = StuMenu{
						MenuName: value.MenuName,
						Icons:    value.Icons,
						Url:      value.Url,
						IsChild:  false,
						Child:    ma,
					}
				} else {
					m := menus[tmpKey]
					m.MenuName, m.Icons, m.Url = value.MenuName, value.Icons, value.Url
					menus[tmpKey] = m
				}
			} else {
				tid := strconv.FormatInt(value.Id, 10)
				if _, vv := menus[tmpKey]; vv {
					mm := menus[tmpKey]
					mm.Child[tid], mm.IsChild = *value, true
					menus[tmpKey] = mm
				} else {
					ma := make(map[string]models.Menu)
					ma[tid] = *value
					menus[tmpKey] = StuMenu{IsChild: true, Child: ma}
				}
			}
		}

		// 判断缓存对象存在存入缓存
		if err == nil {
			if e := oCache.Put("menus", menus, 43200*time.Second); e != nil {
				logs.Alert("存入缓存出现错误 Error:" + e.Error())
			} else {
				logs.Alert("存入缓存数据成功 Success")
			}
		}
	}

	// 使用的布局
	c.Data["admin"] = c.User
	c.Data["navigation"] = menus
	c.Layout = "layout/admin.tpl"
}

// 查询方法
func (c *CommController) Query(search map[string]string) models.Query {
	query := new(models.Query)

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
func (c *CommController) BaseSearch(arr interface{}, search map[string]string, where map[string]interface{}) {
	// 定义返回数据
	var (
		data DataTable
		err  error
	)

	// 处理查询数据信息
	query := c.Query(search)
	for k, v := range where {
		query.Where[k] = v
	}

	// 查询数据
	data.Total, data.Count, err = models.FindAll(arr, query)
	if err != nil {
		c.Error(controllers.CodeBusinessError, "服务器繁忙，请稍后再试", nil)
	}

	data.Data = arr
	c.Success(data, "")
}

// 公共的编辑的方法
func (c *CommController) BaseUpdate(object interface{}, table string) {
	// 获取请求信息
	actionType := c.GetString("actionType")
	if actionType == "" {
		c.Error(controllers.CodeMissingParams, "请求数据为空", nil)
		return
	}

	if !help.InArray([]string{"insert", "update", "delete", "deleteAll"}, actionType) {
		c.Error(controllers.CodeInvalidParams, "请求类型错误", nil)
		return
	}

	// 修改数据
	if actionType == "update" {
		// 修改数据需要先查询数据
		id, err := c.GetInt64("id")
		if err != nil {
			c.Error(controllers.CodeInvalidParams, "主键数据不存在", nil)
			return
		}

		if err := models.One(object, models.QueryOther{Table: table, Where: map[string]interface{}{"id": id}}); err != nil {
			c.Error(controllers.CodeBusinessError, "修改数据不存在", nil)
			return
		}
	}

	if actionType == "deleteAll" {
		ids := c.GetString("ids")
		if ids == "" {
			c.Error(controllers.CodeBusinessError, "删除数据为空", nil)
			return
		}

		aIds := strings.Split(ids, ",")
		if len(aIds) == 0 {
			c.Error(controllers.CodeBusinessError, "删除数据为空", nil)
			return
		}

		if _, err := models.DeleteAll(object, aIds, table); err != nil {
			c.Error(controllers.CodeSystemError, "删除数据失败", nil)
			return
		}

		c.Success(aIds, "批量删除成功")
		return
	}

	if err := c.ParseForm(object); err != nil {
		c.Error(controllers.CodeInvalidParams, "请求数据为空", nil)
		return
	}

	var err error
	// 根据类型做出相应的处理
	switch actionType {
	case "insert": // 新增数据
		_, err = models.Insert(object)
	case "update": // 修改数据
		_, err = models.Update(object)
	case "delete": // 删除数据
		_, err = models.Delete(object)
	}

	// 判断返回数据
	if err != nil {
		c.Error(controllers.CodeBusinessError, "抱歉！执行该操作出现错误", nil)
		return
	}

	c.Success(object, "操作成功")
}

// BaseUpload 图片上传处理
func (c *CommController) BaseUpload(filename, pathname string, allowType []string, size int64, oldFile string) {
	oldPath := c.GetString(oldFile)
	if oldPath != "" {
		// 删除之前的文件
		_ = os.Remove("." + oldPath)
	}

	f, h, err := c.GetFile(filename)
	if err != nil {
		c.Error(controllers.CodeMissingParams, "没有文件上传", nil)
		return
	}

	defer f.Close()
	file := path.Ext(h.Filename)
	if !help.InArray(allowType, file) {
		c.Error(controllers.CodeInvalidParams, "上传文件格式不对", nil)
		return
	}

	if size < f.(help.Sizer).Size() {
		c.Error(controllers.CodeInvalidParams, "上传文件不能超过过大", nil)
		return
	}

	// 处理上传目录
	dirName := "./static/uploads/" + pathname
	// 目录不存在创建
	if !help.IsDirExists(dirName) {
		if err = os.MkdirAll(dirName, 0777); err != nil {
			c.Error(controllers.CodeBusinessError, "创建目录失败 :( "+dirName, nil)
			return
		}
	}

	// 文件最终保存的地址
	fileName := dirName + "/" + strconv.Itoa(rand.Int()) + file
	if err = c.SaveToFile(filename, fileName); err != nil {
		c.Error(controllers.CodeBusinessError, "上传失败", nil)
		return
	}

	data := map[string]string{"fileName": h.Filename, "fileUrl": fileName[1:]}
	c.Success(data, "文件上传成功")
}
