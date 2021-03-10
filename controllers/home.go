package controllers

import (
	"github.com/astaxie/beego/orm"

	"project/models"
)

type HomeController struct {
	Controller
}

// 首页显示
func (i *HomeController) Index() {
	imgList := make([]*models.Image, 0)
	// 查询轮播图片
	orm.NewOrm().QueryTable(&models.Image{}).Filter("status", 1).Filter("type", 1).All(&imgList)
	i.Data["images"] = imgList
	i.Data["action"] = "index"
	i.TplName = "home/index.html"
}

// 获取图片信息
func (i *HomeController) Image() {
	imageList := make([]*models.Image, 0)
	if _, err := orm.NewOrm().QueryTable(&models.Image{}).Filter("status", 1).All(&imageList); err == nil {
		i.Data["images"] = imageList
	}

	i.Data["action"] = "image"
	i.TplName = "home/image.html"
}
