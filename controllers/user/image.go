package user

import (
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"

	"project/connection"
	"project/controllers"
	"project/models"
	"project/response"
	"project/utils"
)

type Image struct {
	controllers.Controller
}

// Create 添加图片信息
func (c *Image) Create() {
	// 未登录
	if !c.IsLogin("user") {
		response.NotLogin(&c.Controller.Controller, "抱歉，您还没有登录呢!")
		return
	}

	// 接收参数
	title, desc, sType, url := c.GetString("title"), c.GetString("desc"), c.GetString("type"), c.GetString("url")
	if title == "" || url == "" {
		response.MissingParams(&c.Controller.Controller, "图片标题和地址为空")
		return
	}

	i, err := strconv.Atoi(sType)
	if err != nil {
		response.InvalidParams(&c.Controller.Controller, "类型错误")
		return
	}

	image := &models.Image{
		Title:       title,
		Description: desc,
		Url:         url,
		Type:        i,
		UserId:      c.User.UserId,
		Status:      1,
	}

	if err := connection.DB.Create(image); err != nil {
		response.SystemError(&c.Controller.Controller, "添加失败")
		return
	}

	response.Success(&c.Controller.Controller, image, "添加成功")
}

// Upload 图片上传
func (c *Image) Upload() {

	// 未登录
	if !c.IsLogin("user") {
		response.NotLogin(&c.Controller.Controller, "抱歉，您还没有登录呢!")
		return
	}

	oldFile := c.GetString("fileurl")

	// 上传文件
	f, h, err := c.GetFile("image")
	if err != nil {
		response.BusinessError(&c.Controller.Controller, err)
		return
	}

	defer f.Close()
	file := path.Ext(h.Filename)
	if !utils.InArray([]string{".jpg", ".jpeg", ".png", ".gif"}, file) {
		response.InvalidParams(&c.Controller.Controller, "上传文件格式不对, 只允许上传: .jpg, .jpeg, .png, .gif文件")
		return
	}

	// 上传文件大小
	if 1024*1024*2 < f.(utils.Sizer).Size() {
		response.InvalidParams(&c.Controller.Controller, "上传文件不能超过2M")
		return
	}

	// 处理上传目录
	datePath := time.Now().Format("200601")
	dirName := "./static/uploads/" + datePath

	// 目录不存在创建
	if !utils.IsDirExists(dirName) {
		if err := os.MkdirAll(dirName, 0777); err != nil {
			response.SystemError(&c.Controller.Controller, "创建目录失败 :( "+dirName)
			return
		}
	}

	// 文件最终保存的地址
	fileName := dirName + "/" + strconv.Itoa(rand.Int()) + file
	if err := c.SaveToFile("image", fileName); err != nil {
		response.SystemError(&c.Controller.Controller, err)
		return
	}

	// 上传成功删除之前的图片
	if oldFile != "" && utils.IsFileExists("./"+oldFile) {
		os.Remove("./" + oldFile)
	}

	data := map[string]string{"image": h.Filename, "path": fileName[1:]}
	response.Success(&c.Controller.Controller, data, "图片上传成功")
	return
}
