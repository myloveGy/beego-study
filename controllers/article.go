package controllers

import (
	"log"

	"github.com/astaxie/beego/orm"

	"project/models"
)

// 首页控制器
type ArticleController struct {
	HomeController
}

func (a *ArticleController) Index() {
	a.Data["action"] = "article"
	a.TplName = "home/article.html"
}

func (a *ArticleController) View() {
	var article, next, prev models.Article
	if err := orm.NewOrm().Raw("SELECT `id`, `title`, `content`, `img`, `create_time`, `see_num`, `comment_num` FROM `my_article` WHERE `status` = ? AND `id` = ? LIMIT 1", 1, a.Ctx.Input.Param(":id")).QueryRow(&article); err != nil {
		log.Println(err)
	}

	// 上一篇
	_ = orm.NewOrm().Raw("SELECT `id`, `title` FROM `my_article` WHERE `status` = ? AND `id` < ? ORDER BY `id` DESC LIMIT 1", 1, article.Id).QueryRow(&prev)

	// 下一篇
	_ = orm.NewOrm().Raw("SELECT `id`, `title` FROM `my_article` WHERE `status` = ? AND `id` > ? ORDER BY `id` ASC LIMIT 1", 1, article.Id).QueryRow(&next)

	a.Data["article"] = article
	a.Data["next"] = next
	a.Data["prev"] = prev
	a.Data["action"] = "article"
	a.TplName = "home/article_view.html"
}

func (a *ArticleController) List() {
	var (
		err         error
		start       int
		length      int
		total       int64
		articleList []models.Article
	)

	// 接收参数
	start, err = a.GetInt("iStart")
	if err != nil {
		a.Error(CodeMissingParams, "参数为空", nil)
		return
	}

	length, err = a.GetInt("iLength")
	if err != nil {
		a.Error(CodeMissingParams, "参数为空", nil)
		return
	}

	m := map[string]interface{}{
		"iTotal":        0,
		"iTotalRecords": 0,
		"aData":         articleList,
	}

	o := orm.NewOrm()
	// 查询数据总条数
	var maps []orm.Params
	if _, err = o.Raw("SELECT COUNT(*) AS `length` FROM `my_article` WHERE `status` = ?", 1).Values(&maps); err != nil {
		a.Error(CodeBusinessError, "查询错误", nil)
		return
	}

	// 查询文章
	total, err = o.Raw("SELECT `id`, `title`, `content`, `img`, `create_time`, `see_num`, `comment_num` FROM `my_article` WHERE `status` = ? LIMIT ?, ?", 1, start, length).QueryRows(&articleList)
	if err != nil {
		a.Error(CodeBusinessError, "查询错误", nil)
		return
	}

	m["iTotal"] = maps[0]["length"]
	m["iTotalRecords"] = total
	m["aData"] = articleList

	a.Success(m, "success")
}
