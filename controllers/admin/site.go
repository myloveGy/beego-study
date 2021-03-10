package admin

import (
	"crypto/sha1"
	"fmt"
	"io"

	"github.com/astaxie/beego/orm"

	"project/controllers"
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

	// 密码加密
	h := sha1.New()
	io.WriteString(h, password)
	password = fmt.Sprintf("%x", h.Sum(nil))

	var u controllers.User
	// 查询数据
	if err := orm.NewOrm().Raw(
		"SELECT `id`, `username`, `email`, `status` FROM `my_admin` WHERE `username` = ? AND `password` = ? LIMIT 1",
		username,
		password).QueryRow(&u); err != nil {
		s.Error(controllers.CodeInvalidParams, "用户不存在或者密码错误", nil)
		return
	}

	if u.Status != 1 {
		s.Error(controllers.CodeBusinessError, "对不起！你被管理员封好了 ):", nil)
		return
	}

	// 设置session
	s.SetSession("admin", u)
	s.User = u
	s.Success(map[string]interface{}{
		"username": username,
		"email":    u.Email,
		"user_id":  u.Id,
	}, "登录成功")
}

// 用户退出
func (s *SiteController) Logout() {
	s.DelSession("admin")
	s.Success(nil, "退出成功")
}
