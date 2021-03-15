package user

import (
	"github.com/astaxie/beego/orm"

	"project/controllers"
	"project/models"
	"project/response"
)

type Article struct {
	controllers.Controller
}

// Create 新增文章信息
func (a *Article) Create() {
	// 未登录
	if !a.IsLogin("user") {
		response.NotLogin(&a.Controller.Controller, "抱歉，您还没有登录呢!")
		return
	}

	var article models.Article
	if err := a.ParseForm(&article); err != nil {
		response.MissingParams(&a.Controller.Controller, "提交数据为空: "+err.Error())
		return
	}

	// 附加信息
	article.UserId = a.User.UserId
	article.Status = 1
	if _, err := orm.NewOrm().Insert(&article); err != nil {
		response.SystemError(&a.Controller.Controller, "添加文章失败")
		return
	}

	response.Success(&a.Controller.Controller, &article, "添加成功")
}
