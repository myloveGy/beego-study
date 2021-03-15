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
func (s *Guest) Index() {

	// 用户已经登录
	if s.IsLogin("admin") {
		s.Redirect("/admin/admin/site", 302)
	}

	s.TplName = "admin/login.html"
}

// 执行登录
func (s *Guest) Login() {
	// 获取参数
	username, password := s.GetString("username"), s.GetString("password")
	if username == "" && password == "" {
		response.MissingParams(&s.Base.Controller, "参数为空")
		return
	}

	// 用户登录
	admin := &models.Admin{}
	if err := admin.Login(username, password, s.Ctx.Request.RemoteAddr); err != nil {
		response.BusinessError(&s.Base.Controller, err)
		return
	}

	s.User = response.User{UserId: admin.UserId, Username: admin.Username, Status: admin.Status, Email: admin.Email}

	// 设置session
	s.SetSession("admin", s.User)
	response.Success(&s.Base.Controller, &s.User, "登录成功")
}

// 用户退出
func (s *Guest) Logout() {
	cache.Delete("menu")
	s.DelSession("admin")
	s.Redirect("/admin", 302)
}
