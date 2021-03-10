package user

import (
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"

	"project/controllers"
	"project/help"
	"project/models"
)

type ImageController struct {
	controllers.HomeController
}

// Create 添加图片信息
func (c *ImageController) Create() {
	// 未登录
	if !c.IsLogin("user") {
		c.Error(controllers.CodeNotLogin, "抱歉，您还没有登录呢!", nil)
		return
	}

	// 接收参数
	title, desc, sType, url := c.GetString("title"), c.GetString("desc"), c.GetString("type"), c.GetString("url")
	if title == "" || url == "" {
		c.Error(controllers.CodeMissingParams, "图片标题和地址为空", nil)
		return
	}

	i, err := strconv.Atoi(sType)
	if err != nil {
		c.Error(controllers.CodeInvalidParams, "类型错误", nil)
		return
	}

	image := &models.Image{
		Title:       title,
		Description: desc,
		Url:         url,
		Type:        i,
		UserId:      c.User.Id,
	}
	if _, err := orm.NewOrm().Insert(image); err != nil {
		c.Error(controllers.CodeSystemError, "添加失败", nil)
		return
	}

	c.Success(image, "添加成功")
}

// Upload 图片上传
func (c *ImageController) Upload() {
	oldFile := c.GetString("fileurl")

	// 未登录
	if !c.IsLogin("user") {
		c.Error(controllers.CodeNotLogin, "抱歉，您还没有登录呢!", nil)
		return
	}

	// 上传文件
	f, h, err := c.GetFile("image")
	if err != nil {
		c.Error(controllers.CodeBusinessError, err.Error(), nil)
		return
	}

	defer f.Close()
	file := path.Ext(h.Filename)
	if !help.InArray([]string{".jpg", ".jpeg", ".png", ".gif"}, file) {
		c.Error(controllers.CodeInvalidParams, "上传文件格式不对, 只允许上传: .jpg, .jpeg, .png, .gif文件", nil)
		return
	}

	// 上传文件大小
	if 1024*1024*2 < f.(help.Sizer).Size() {
		c.Error(controllers.CodeInvalidParams, "上传文件不能超过2M", nil)
		return
	}

	// 处理上传目录
	datePath := time.Now().Format("200601")
	dirName := "./static/uploads/" + datePath

	// 目录不存在创建
	if !help.IsDirExists(dirName) {
		if err := os.MkdirAll(dirName, 0777); err != nil {
			c.Error(controllers.CodeSystemError, "创建目录失败 :( "+dirName, nil)
			return
		}
	}

	// 文件最终保存的地址
	fileName := dirName + "/" + strconv.Itoa(rand.Int()) + file
	if err := c.SaveToFile("image", fileName); err != nil {
		c.Error(controllers.CodeSystemError, err.Error(), nil)
		return
	}

	// 上传成功删除之前的图片
	if oldFile != "" && help.IsFileExists("./"+oldFile) {
		os.Remove("./" + oldFile)
	}

	data := map[string]string{"image": h.Filename, "path": fileName[1:]}
	c.Success(data, "图片上传成功")
	return
}
