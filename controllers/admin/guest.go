package admin

import (
	"project/cache"
	"project/controllers"
	"project/models"
	"project/response"
)

// 后台首页控制器
type Guest struct {
	controllers.Base
}

// 显示登录页面
func (g *Guest) Index() {

	// 用户已经登录
	if g.IsLogin("admin") {
		g.Redirect("/admin/admin/site", 302)
	}

	g.TplName = "admin/login.html"
}

// 执行登录
func (g *Guest) Login() {
	// 获取参数
	username, password := g.GetString("username"), g.GetString("password")
	if username == "" && password == "" {
		response.MissingParams(&g.Base.Controller, "参数为空")
		return
	}

	// 用户登录
	admin := &models.Admin{}
	if err := admin.Login(username, password, g.Ctx.Request.RemoteAddr); err != nil {
		response.BusinessError(&g.Base.Controller, err)
		return
	}

	g.User = response.User{UserId: admin.UserId, Username: admin.Username, Status: admin.Status, Email: admin.Email}

	// 设置session
	g.SetSession("admin", g.User)
	response.Success(&g.Base.Controller, &g.User, "登录成功")
}

// 用户退出
func (g *Guest) Logout() {
	cache.Delete("menu")
	g.DelSession("admin")
	g.Redirect("/admin", 302)
}
