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

// 获取文件大小的接口
type Sizer interface {
	Size() int64
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

// 判断文件是否存在
func isFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 判断文件文件类型
func AllowType(StrFile string, AllType []string) (isAllow bool) {
	for _, v := range AllType {
		if StrFile == v {
			isAllow = true
			return
		}
	}

	return
}

func (this *ImageController) Index() {
	this.TplNames = "Admin/image.html"
}

// 图片查看
func (this *ImageController) View() {
	// 查询图片信息
	query := map[string]interface{}{"status": 1}
	image, err := models.GetAllImage(query)
	if err == nil {
		this.Data["images"] = image
	}
	this.TplNames = "Admin/image_view.html"
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
	Point, oldFile := this.InitPoint(), this.GetString("Url")
	f, h, err := this.GetFile("fileUrl")
	defer f.Close()

	// 判断错误
	if err == nil {

		// 获取上传文件的大小和文件类型
		lastNum := strings.LastIndex(h.Filename, ".")
		IntSize, StrFile := f.(Sizer).Size(), h.Filename[lastNum:]

		// 定义允许上传
		Point.Message = "上传文件类型错误"
		allType := []string{".jpeg", ".jpg", ".png", ".gif"}
		if AllowType(StrFile, allType) {

			// 判断上传文件大小
			Point.Message = "上传文件大小超过2MB"
			if IntSize < 1024*1024*2 {

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

					// 文件最终保存的地址
					fileName := dirName + "/" + strconv.Itoa(RandName) + StrFile
					err = this.SaveToFile("fileUrl", fileName)

					if err == nil {
						// 上传成功删除之前的图片
						if oldFile != "" && isFileExists("./"+oldFile) {
							os.Remove("./" + oldFile)
						}

						Point.Status = 1
						Point.Message = "图片上传成功"
						Point.Data = fileName[1:]
					} else {
						Point.Message = err.Error()
					}
				}
			}
		}

	} else {
		Point.Message = err.Error()
	}

	this.AjaxReturn(Point)
}
