package controllers

import (
	"project/connection"
	"project/models"
)

type Home struct {
	Controller
}

// 首页显示
func (i *Home) Index() {
	// 查询轮播图片
	imgList := make([]*models.Image, 0)
	connection.DB.Builder(&imgList).Where("status", 1).Where("type", 1).All()
	i.Data["images"] = imgList
	i.Data["action"] = "index"
	i.TplName = "home/index.html"
}
