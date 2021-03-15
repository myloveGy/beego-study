package controllers

import (
	"github.com/astaxie/beego"

	"project/connection"
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

	var (
		sees     = make([]*models.Article, 0)
		hots     = make([]*models.Article, 0)
		commList = make([]*models.Article, 0)
		imgList  = make([]*models.Article, 0)
	)

	// 点击量
	connection.DB.FindAll(&sees, "`status` = ? ORDER BY `see_num` desc LIMIT 6", 1)
	// 热门
	connection.DB.FindAll(&hots, "`status` = ? ORDER BY `created_at` desc LIMIT 6", 1)
	// 推荐
	connection.DB.FindAll(&commList, "`status` = ? AND `recommend` = ?  ORDER BY `created_at` desc LIMIT 6", 1, 1)
	// 图文推荐
	connection.DB.FindAll(&imgList, "`status` = ? AND `img` != ? ORDER BY `created_at` asc LIMIT 5", 1, "")

	c.Layout = "layout/home.html"
	c.Data["isLogin"] = c.IsLogin("user")
	c.Data["sees"] = sees
	c.Data["hots"] = hots
	c.Data["commList"] = commList
	c.Data["imgList"] = imgList
	c.Data["user"] = c.User.Username
	c.Data["me"], _ = beego.AppConfig.GetSection("me")
}
