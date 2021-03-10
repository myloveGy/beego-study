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
	hots, _ := models.GetArticle(status, 6, "-created_at")
	status["recommend"] = 1

	// 推荐
	commList, _ := models.GetArticle(status, 6, "-created_at")

	// 图文推荐
	var imgList []*models.Article
	orm.NewOrm().QueryTable(&models.Article{}).Filter("status", 1).Exclude("img", "").OrderBy("created_at").Limit(5).All(&imgList)

	// 用户是否已经登录
	this.Data["isLogin"] = this.IsLogin("user")
	this.Data["sees"] = sees
	this.Data["hots"] = hots
	this.Data["commList"] = commList
	this.Data["imgList"] = imgList
	this.Data["user"] = this.User.Username
}
