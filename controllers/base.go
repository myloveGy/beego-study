package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"project/models"
	"project/response"
)

// 继承基础控制器
type Base struct {
	beego.Controller
	User response.User
}

// 验证用户是否已经登录
func (b *Base) IsLogin(str string) bool {
	user := b.GetSession(str)
	if user == nil {
		return false
	}

	b.User = user.(response.User)
	return b.User.UserId > 0
}

// 首页控制器
type Controller struct {
	Base
}

func (c *Controller) Prepare() {

	if c.IsAjax() {
		return
	}

	// 使用的布局
	c.Layout = "layout/home.html"
	status := map[string]interface{}{"status": 1}

	// 点击量
	sees, _ := models.GetArticle(status, 6, "-see_num")

	// 热门
	hots, _ := models.GetArticle(status, 6, "-created_at")
	status["recommend"] = 1

	// 推荐
	commList, _ := models.GetArticle(status, 6, "-created_at")

	// 图文推荐
	var imgList []*models.Article
	orm.NewOrm().QueryTable(&models.Article{}).
		Filter("status", 1).
		Exclude("img", "").
		OrderBy("created_at").
		Limit(5).
		All(&imgList)

	// 用户是否已经登录
	c.Data["isLogin"] = c.IsLogin("user")
	c.Data["sees"] = sees
	c.Data["hots"] = hots
	c.Data["commList"] = commList
	c.Data["imgList"] = imgList
	c.Data["user"] = c.User.Username
}
