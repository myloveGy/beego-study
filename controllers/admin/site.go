package admin

import (
	"project/controllers"
	"project/models"
)

// 后台首页控制器
type SiteController struct {
	controllers.BaseController
}

// 显示登录页面
func (s *SiteController) Index() {
	// 用户已经登录
	if s.IsLogin("admin") {
		s.Redirect("/admin/site", 302)
	}

	s.TplName = "layout/login.html"
}

// 执行登录
func (s *SiteController) Login() {

	// 获取参数
	username, password := s.GetString("username"), s.GetString("password")
	if username == "" && password == "" {
		s.Error(controllers.CodeMissingParams, "参数为空", nil)
		return
	}

	// 用户登录
	admin := &models.Admin{}
	if err := admin.Login(username, password, s.Ctx.Request.RemoteAddr); err != nil {
		s.Error(controllers.CodeBusinessError, err.Error(), nil)
		return
	}

	s.User = controllers.User{UserId: admin.UserId, Username: admin.Username, Status: admin.Status, Email: admin.Email}

	// 设置session
	s.SetSession("admin", s.User)
	s.Success(&s.User, "登录成功")
}

// 用户退出
func (s *SiteController) Logout() {
	s.DelSession("admin")
	s.Success(nil, "退出成功")
}
