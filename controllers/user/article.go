package user

import (
	"github.com/astaxie/beego/orm"

	"project/controllers"
	"project/models"
)

type ArticleController struct {
	controllers.HomeController
}

// Create 新增文章信息
func (a *ArticleController) Create() {
	// 未登录
	if !a.IsLogin("user") {
		a.Error(controllers.CodeNotLogin, "抱歉，您还没有登录呢!", nil)
		return
	}

	var article models.Article
	if err := a.ParseForm(&article); err != nil {
		a.Error(controllers.CodeMissingParams, "提交数据为空: "+err.Error(), nil)
		return
	}

	// 附加信息
	article.UserId = a.User.Id
	article.Status = 1
	if _, err := orm.NewOrm().Insert(&article); err != nil {
		a.Error(controllers.CodeSystemError, "添加文章失败", nil)
		return
	}

	a.Success(article, "添加成功")
}
