package controllers

import (
	"log"

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
		log.Println("see-num", err)
	}

	// 热门
	hots, err1 := models.GetArticle(status, 6, "-create_time")
	if err1 != nil {
		log.Println("create_time", err)
	}

	status["recommend"] = 1

	// 推荐
	commList, err2 := models.GetArticle(status, 6, "-create_time")
	if err2 != nil {
		log.Println("create_time 1", err)
	}

	// 图文推荐
	var imgList []models.Article
	_, err = orm.NewOrm().Raw("SELECT `id`, `title`, `img`, `create_time` FROM `my_article` WHERE `status` = ? AND `img` != '' ORDER BY `create_time` LIMIT 5", 1).QueryRows(&imgList)
	if err != nil {
		log.Println("create_time 2", err)
	}

	// 用户是否已经登录
	this.Data["isLogin"] = this.IsLogin("user")
	this.Data["sees"] = sees
	this.Data["hots"] = hots
	this.Data["commList"] = commList
	this.Data["imgList"] = imgList
	this.Data["user"] = this.User.Username
}
