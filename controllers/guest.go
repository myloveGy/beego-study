package controllers

import (
	"project/models"
)

type Guest struct {
	Base
}

// Login 登录
func (g *Guest) Login() {
	// 获取参数
	username, password := g.GetString("username"), g.GetString("password")
	if username == "" && password == "" {
		g.Error(CodeMissingParams, "参数为空", nil)
		return
	}

	// 用户登录
	admin := &models.Admin{}
	if err := admin.Login(username, password, g.Ctx.Request.RemoteAddr); err != nil {
		g.Error(CodeBusinessError, err.Error(), nil)
		return
	}

	g.User = User{UserId: admin.UserId, Username: admin.Username, Status: admin.Status, Email: admin.Email}

	// 设置session
	g.SetSession("user", g.User)
	g.Success(&g.User, "登录成功")
}

// Logout 退出
func (g *Guest) Logout() {
	// 初始化返回
	g.DelSession("user")
	g.Success(nil, "您已经退出登录")
}

// Detail 详情
func (g *Guest) Detail() {
	g.Success(g.GetSession("user"), "")
}
