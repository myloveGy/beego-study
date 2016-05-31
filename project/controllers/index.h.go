package controllers

import (
	"github.com/astaxie/beego/orm"
	"project/models"
)

// 首页控制器
type IndexHomeController struct {
	HomeController
}

// 定义图片类型
type Image struct {
	Title string
	Desc  string
	Url   string
}

// 首页显示
func (this *IndexHomeController) Get() {
	var imgs []Image
	// 查询轮播图片
	_, err := orm.NewOrm().Raw("SELECT `title`, `desc`, `url` FROM `my_image` WHERE `type` = ? AND `status` = ?", 1, 1).QueryRows(&imgs)
	if err != nil {
		this.Redirect("/", 302)
	}

	// 查询文章
	var articls []models.Article
	_, err1 := orm.NewOrm().Raw("SELECT `id`, `title`, `content`, `img`, `create_time`, `see_num`, `comment_num` FROM `my_article` WHERE `status` = ? AND `img` != ?", 1, "").QueryRows(&articls)
	if err1 != nil {
		this.Redirect("/", 302)
	}
	this.Data["images"] = imgs
	this.Data["articles"] = articls
	this.Data["action"] = "index"
	this.TplName = "home/index.html"
}

// 请求获取文章信息
func (this *IndexHomeController) Ajax() {
	this.E = ArrError{Status: 1, Message: "参数为空", Data: nil}
	this.AjaxReturn()
}
