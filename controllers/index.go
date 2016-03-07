package controllers

import (
	models "app/models"
	"github.com/astaxie/beego"
	"time"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "Layout/login.html"
}

// 用户登录
func (c *IndexController) Login() {

	// 初始化定义错误
	ArrData := struct {
		Status  int
		Message string
		Data    interface{}
	}{
		Status:  0,
		Message: "参数为空",
		Data:    "",
	}

	// 接收参数
	admin := c.GetString("username")
	pass := c.GetString("password")

	// 验证登录
	if admin != "" && pass != "" {
		ArrData.Message = "用户名或者密码错误"
		IsTrue, objAdmin := models.AdminGetOne(admin, pass)
		if IsTrue == true {
			c.SetSession("AdminUser", objAdmin)
			ArrData.Status = 1
			ArrData.Message = "登录成功,正在进行页面跳转,请稍候..."
		}
	}

	c.Data["json"] = &ArrData
	c.ServeJson()
}

// 用户退出
func (this *IndexController) Logout() {
	// 修改用户信息
	admin := this.GetSession("AdminUser")
	if admin != nil {
		adminUser := admin.(models.Admin)
		// 执行修改
		_ = models.AdminLogout(adminUser.Id, this.Ctx.Request.RemoteAddr, time.Now().Unix())
	}

	// 清除SESSION
	this.DelSession("AdminUser")
	this.Redirect("/", 302)
}
