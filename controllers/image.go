package controllers

import (
	models "app/models"
	// "github.com/astaxie/beego"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type ImageController struct {
	BaseController
}

//  判断目录是否存在
func isDirExists(path string) bool {
	file, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return file.IsDir()
	}

	panic("no reached")
}

func (this *ImageController) Index() {
	this.TplNames = "Admin/image.html"
}

// 响应ajax获取数据
func (this *ImageController) AjaxIndex() {
	// 默认返回
	Point := this.InitPoint()

	// 定义查询map类型
	mMap := map[string]string{
		"Status": "status",
		"Type":   "type",
		"search": "title__contains",
		"Id":     "id",
	}

	// 查询字符串
	tmpMap, offset, limit, order := this.GetQueryString(mMap, "id")

	if this.IsAjax() {
		Point.Message = "获取数据为空"
		total, count, data, err := models.ImageGetAll(tmpMap, offset, limit, order)
		if err == nil {
			Point.Status = 1
			Point.Message = "success"
			Point.Data = this.DataTable(total, count, data)
		}

	}

	this.AjaxReturn(Point)
}

// 响应其他操作
func (this *ImageController) Save() {
	// 定义错误
	Point := this.InitPoint()

	// 接收参数
	actionType, IsHave, Image := this.GetString("actionType"), false, models.Image{}
	var Id, RowCount int64

	// 判断提交数据
	if actionType != "" {
		Point.Message = "数据赋值失败..."
		if err := this.ParseForm(&Image); err == nil {
			// 获取SESSION值
			admin := this.GetSession("AdminUser")

			// 获取修改用户ID
			if admin != nil {
				user := admin.(models.Admin)
				Image.UpdateId = user.Id
			}

			// 判断类型操作
			switch actionType {
			case "delete":
				IsHave, RowCount, err = models.ImageDelete(Image.Id)
			case "insert":
				IsHave = true
				Id, err = models.ImageInsert(&Image)
			case "update":
				IsHave, RowCount, err = models.ImageSave(&Image)
			}

			// 判断操作是否成功
			Point.Message = "抱歉!没有查询到数据..."
			if IsHave {
				Point.Message = "抱歉！服务器繁忙,请稍候再试..."
				if RowCount > 0 || Id > 0 {
					Point.Status = 1
					Point.Message = "恭喜你！操作成功 ^.^ "
				}
			}
		}
	}

	this.AjaxReturn(Point)
}

// 图片上传
func (this *ImageController) FileUpload() {
	// 定义错误
	Point := this.InitPoint()
	f, h, err := this.GetFile("fileUrl")
	defer f.Close()

	if err == nil {
		// 判断上传文件类型

		// 判断上传文件大小

		// 处理上传目录
		datePath := time.Now().Format("200601")
		dirName := "./static/img/" + datePath

		// 目录不存在创建
		if !isDirExists(dirName) {
			err = os.MkdirAll(dirName, 0777)
		}

		Point.Message = "创建目录失败 :( " + dirName

		// 创建目录失败
		if err == nil {

			// 处理上传的文件名
			RandName := rand.Int()
			lastNum := strings.LastIndex(h.Filename, ".")
			file := h.Filename[lastNum:]

			// 文件最终保存的地址
			fileName := dirName + "/" + strconv.Itoa(RandName) + file
			err = this.SaveToFile("fileUrl", fileName)

			if err == nil {
				Point.Status = 1
				Point.Message = "图片上传成功"
				Point.Data = fileName
			} else {
				Point.Message = err.Error()
			}
		}

	} else {
		Point.Message = err.Error()
		Point.Message = "123"
	}

	this.AjaxReturn(Point)
}
