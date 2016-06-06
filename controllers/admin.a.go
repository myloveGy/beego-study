package controllers

import (
	"math/rand"
	"os"
	"path"
	"strconv"
)

type AdminAdminController struct {
	AdminController
}

// @router /admin/site:k [get]
func (this *AdminAdminController) Site() {
	this.Data["image"] = []int{1, 2, 3, 4, 5, 6, 1, 2}
	this.TplName = "admin/site.html"
}

// 修改编辑
func (this *AdminAdminController) Update() {
	this.E = ArrError{Status: 0, Msg: "参数为空", Data: nil}
	this.AjaxReturn()
}

// 图片上传
func (this *AdminAdminController) Upload() {
	this.E = ArrError{Status: 0, Msg: "抱歉，您还没有登录呢!", Data: nil}
	if this.isLogin("admin") {
		f, h, err := this.GetFile("avatar")
		defer f.Close()
		if err == nil {
			file := path.Ext(h.Filename)
			this.E.Msg = "上传文件格式不对"
			if InArray([]string{".jpg", ".jpeg", ".png", ".gif"}, file) {
				this.E.Msg = "上传文件不能超过2M"
				if 1024*1024*2 > f.(Sizer).Size() {
					// 处理上传目录
					dirName := "./static/uploads/avatar"

					// 目录不存在创建
					if !isDirExists(dirName) {
						err = os.MkdirAll(dirName, 0777)
					}

					this.E.Msg = "创建目录失败 :( " + dirName

					// 创建目录失败
					if err == nil {
						// 文件最终保存的地址
						fileName := dirName + "/" + strconv.Itoa(rand.Int()) + file
						err = this.SaveToFile("avatar", fileName)
						if err == nil {
							this.E.Status = 1
							this.E.Msg = "图片上传成功"
							this.E.Data = map[string]string{"image": h.Filename, "fileDir": fileName[1:]}
						} else {
							this.E.Msg = err.Error()
						}
					}

				}
			}
		} else {
			this.E.Msg = err.Error()
		}
	}

	this.AjaxReturn()
}
