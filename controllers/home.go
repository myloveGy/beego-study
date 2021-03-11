package controllers

import (
	"github.com/astaxie/beego/orm"

	"project/models"
)

type Home struct {
	Controller
}

// 首页显示
func (i *Home) Index() {
	imgList := make([]*models.Image, 0)
	// 查询轮播图片
	orm.NewOrm().QueryTable(&models.Image{}).Filter("status", 1).Filter("type", 1).All(&imgList)
	i.Data["images"] = imgList
	i.Data["action"] = "index"
	i.TplName = "home/index.html"
}
