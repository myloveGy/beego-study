package controllers

import (
	"github.com/astaxie/beego/orm"

	"project/models"
)

type ImageController struct {
	Controller
}

// Index 获取图片信息
func (i *ImageController) Index() {
	imageList := make([]*models.Image, 0)
	if _, err := orm.NewOrm().QueryTable(&models.Image{}).Filter("status", 1).All(&imageList); err == nil {
		i.Data["images"] = imageList
	}

	i.Data["action"] = "image"
}
