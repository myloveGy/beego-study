package controllers

import (
	"github.com/astaxie/beego/orm"
	"project/models"
)

// 首页控制器
type HomeController struct {
	BaseController
}

func (this *HomeController) Prepare() {
	// 使用的布局
	this.Layout = "layout/home.tpl"
	status := map[string]interface{}{"status": 1}

	// 点击量
	sees, err := models.GetArticle(status, 6, "-see_num")
	if err != nil {
		this.Redirect("/", 302)
	}

	// 热门
	hots, err1 := models.GetArticle(status, 6, "-create_time")
	if err1 != nil {
		this.Redirect("/", 302)
	}

	status["recommend"] = 1
	// 推荐
	comms, err2 := models.GetArticle(status, 6, "-create_time")
	if err2 != nil {
		this.Redirect("/", 302)
	}

	// 图文推荐
	var imgs []models.Article
	_, err = orm.NewOrm().Raw("SELECT `id`, `title`, `img`, `create_time` FROM `my_article` WHERE `status` = ? AND `img` != '' ORDER BY `create_time` LIMIT 5", 1).QueryRows(&imgs)
	if err != nil {
		this.Redirect("/", 302)
	}

	// 用户是否已经登录
	this.Data["isLogin"] = this.isLogin("user")
	this.Data["sees"] = sees
	this.Data["hots"] = hots
	this.Data["comms"] = comms
	this.Data["imgs"] = imgs
	this.Data["user"] = this.U.Username
}
