package controllers

import (
	"project/connection"
	"project/models"
)

type Image struct {
	Controller
}

// Index 获取图片信息
func (i *Image) Index() {
	imageList := make([]*models.Image, 0)
	if err := connection.DB.Builder(&imageList).Where("status", 1).All(); err == nil {
		i.Data["images"] = imageList
	}

	i.Data["action"] = "image"
	i.TplName = "image/index.html"
}
