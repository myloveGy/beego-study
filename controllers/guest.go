package controllers

import (
	"project/models"
	"project/response"
)

type Guest struct {
	Base
}

// Login 登录
func (g *Guest) Login() {
	// 获取参数
	username, password := g.GetString("username"), g.GetString("password")
	if username == "" && password == "" {
		response.MissingParams(&g.Controller)
		return
	}

	// 用户登录
	admin := &models.Admin{}
	if err := admin.Login(username, password, g.Ctx.Request.RemoteAddr); err != nil {
		response.MissingParams(&g.Controller, err)
		return
	}

	g.User = response.User{UserId: admin.UserId, Username: admin.Username, Status: admin.Status, Email: admin.Email}

	// 设置session
	g.SetSession("user", g.User)
	response.Success(&g.Controller, &g.User, "登录成功")
	return
}

// Logout 退出
func (g *Guest) Logout() {
	// 初始化返回
	g.DelSession("user")
	g.Redirect("/", 302)
}

// Detail 详情
func (g *Guest) Detail() {
	response.Success(&g.Controller, g.GetSession("user"), "")
}
