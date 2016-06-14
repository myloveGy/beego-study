package controllers

import (
	"github.com/astaxie/beego/orm"
	"project/models"
)

// 首页控制器
type ArticleController struct {
	HomeController
}

func (this *ArticleController) Index() {
	this.Data["action"] = "article"
	this.TplName = "home/article.html"
}

func (this *ArticleController) View() {
	var article, next, prev models.Article
	if orm.NewOrm().Raw("SELECT `id`, `title`, `content`, `img`, `create_time`, `see_num`, `comment_num` FROM `my_article` WHERE `status` = ? AND `id` = ? LIMIT 1", 1, this.Ctx.Input.Param(":id")).QueryRow(&article) != nil {
		this.Redirect("/", 302)
	}

	// 上一篇
	_ = orm.NewOrm().Raw("SELECT `id`, `title` FROM `my_article` WHERE `status` = ? AND `id` < ? ORDER BY `id` DESC LIMIT 1", 1, article.Id).QueryRow(&prev)

	// 下一篇
	_ = orm.NewOrm().Raw("SELECT `id`, `title` FROM `my_article` WHERE `status` = ? AND `id` > ? ORDER BY `id` ASC LIMIT 1", 1, article.Id).QueryRow(&next)

	this.Data["article"] = article
	this.Data["next"] = next
	this.Data["prev"] = prev
	this.Data["action"] = "article"
	this.TplName = "home/article_view.html"
}

func (this *ArticleController) List() {
	// 初始化返回
	this.E = ArrError{Status: 0, Msg: "参数为空", Data: nil}
	var iStart, iLength int
	var err error
	var articls []models.Article
	// 接收参数
	iStart, err = this.GetInt("iStart")
	if err == nil {
		iLength, err = this.GetInt("iLength")
		if err == nil {
			m := map[string]interface{}{"iTotal": 0, "iTotalRecords": 0, "aData": articls}
			o := orm.NewOrm()
			// 查询数据总条数
			var maps []orm.Params
			_, err = o.Raw("SELECT COUNT(*) AS `length` FROM `my_article` WHERE `status` = ?", 1).Values(&maps)
			if err == nil {
				// 查询文章
				num, err1 := o.Raw("SELECT `id`, `title`, `content`, `img`, `create_time`, `see_num`, `comment_num` FROM `my_article` WHERE `status` = ? LIMIT ?, ?", 1, iStart, iLength).QueryRows(&articls)
				if err1 == nil {
					m["iTotal"] = maps[0]["length"]
					m["iTotalRecords"] = num
					m["aData"] = articls
					this.E.Status = 1
					this.E.Data = m
					this.AjaxReturn()
				} else {
					this.E.Msg = err1.Error()
				}
			} else {
				this.E.Msg = err.Error()
			}

		} else {
			this.E.Msg = err.Error()
		}
	} else {
		this.E.Msg = err.Error()
	}

	this.AjaxReturn()
}
