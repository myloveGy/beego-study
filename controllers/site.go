package controllers

import (
	"crypto/sha1"
	"fmt"
	"io"

	"github.com/astaxie/beego/orm"
)

// 后台首页控制器
type SiteController struct {
	BaseController
}

// 显示登录页面
func (this *SiteController) Index() {
	// 用户已经登录
	if this.isLogin("admin") {
		this.Redirect("/admin/site", 302)
	}
	this.TplName = "layout/login.html"
}

// 执行登录
func (this *SiteController) Login() {
	// 初始化返回
	this.E = ArrError{Status: 0, Msg: "参数为空", Data: nil}
	// 获取参数
	username, password := this.GetString("username"), this.GetString("password")
	if username != "" && password != "" {
		// 密码加密
		h := sha1.New()
		io.WriteString(h, password)
		password = fmt.Sprintf("%x", h.Sum(nil))

		var u User
		// 查询数据
		err := orm.NewOrm().Raw("SELECT `id`, `username`, `email`, `status` FROM `my_admin` WHERE `username` = ? AND `password` = ? LIMIT 1", username, password).QueryRow(&u)
		this.E.Msg = "用户不存在或者密码错误"
		if err == nil {
			this.E.Msg = "对不起！你被管理员封号了 ):"
			if u.Status == 1 {
				// 设置session
				this.SetSession("admin", u)
				this.U = u
				this.E.Status = 1
				this.E.Msg = "登录成功 ^.^ !"
				this.E.Data = u.Username
			}
		}
	}

	this.AjaxReturn()
}

// 用户退出
func (this *SiteController) Logout() {
	this.E = ArrError{Status: 0, Msg: "您已经退出登录", Data: nil}
	if this.GetSession("admin") != nil {
		this.DelSession("admin")
		this.E.Status = 1
	}
	this.AjaxReturn()
}
