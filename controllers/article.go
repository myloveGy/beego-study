package controllers

import (
	"log"
	"strconv"

	"github.com/astaxie/beego/orm"

	"project/models"
	"project/response"
)

// 首页控制器
type Article struct {
	Controller
}

// Index 文章列表
func (a *Article) Index() {
	a.Data["action"] = "article"
	a.TplName = "article/index.html"
}

// View 文章详情
func (a *Article) Detail() {
	var (
		article, next, prev models.Article
		o                   = orm.NewOrm()
	)

	a.Data["action"] = "article"
	a.TplName = "article/detail.html"
	id := a.Ctx.Input.Param("0")
	if id == "" {
		return
	}

	article.Id, _ = strconv.ParseInt(id, 10, 64)
	if article.Id > 0 {
		if err := o.Read(&article); err != nil {
			log.Println(err)
		}
	}

	// 上一篇
	_ = o.QueryTable(&prev).
		Filter("status", 1).
		Filter("id__lt", article.Id).
		OrderBy("-id").
		Limit(1).
		One(&prev)

	// 下一篇
	_ = o.QueryTable(&next).
		Filter("status", 1).
		Filter("id__gt", article.Id).
		OrderBy("id").Limit(1).
		One(&next)

	a.Data["article"] = article
	a.Data["next"] = next
	a.Data["prev"] = prev
}

// List 文章列表
func (a *Article) List() {
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
		response.MissingParams(&a.Base.Controller)
		return
	}

	length, err = a.GetInt("iLength")
	if err != nil {
		response.MissingParams(&a.Base.Controller)
		return
	}

	m := map[string]interface{}{
		"iTotal":        0,
		"iTotalRecords": 0,
		"aData":         articleList,
	}

	o := orm.NewOrm()

	// 查询数据总条数
	total, err = o.QueryTable(&models.Article{}).Filter("status", 1).Count()
	if err != nil {
		response.BusinessError(&a.Base.Controller, "查询错误")
		return
	}

	if _, err := o.QueryTable(&models.Article{}).Filter("status", 1).Limit(length, start).All(&articleList); err != nil {
		response.BusinessError(&a.Base.Controller, "查询错误")
		return
	}

	m["iTotal"] = total
	m["iTotalRecords"] = len(articleList)
	m["aData"] = articleList

	response.Success(&a.Base.Controller, m)
}

// Image 请求获取图片文章信息
func (a *Article) Image() {
	// 初始化返回
	var (
		start, length int
		total         int64
		err           error
	)

	// 接收参数
	start, err = a.GetInt("iStart")
	if err != nil {
		response.MissingParams(&a.Base.Controller)
		return
	}

	length, err = a.GetInt("iLength")
	if err != nil {
		response.MissingParams(&a.Base.Controller)
		return
	}

	o := orm.NewOrm()
	// 查询数据总条数
	total, err = o.QueryTable(&models.Article{}).Filter("status", 1).FilterRaw("img", "!= ''").Count()
	if err != nil {
		response.BusinessError(&a.Base.Controller, "查询数据为空")
		return
	}

	// 查询文章
	articleList := make([]*models.Article, 0)
	if _, err := o.QueryTable(&models.Article{}).
		Filter("status", 1).
		FilterRaw("img", "!= ''").
		OrderBy("-id").
		Limit(length, start).
		All(&articleList); err != nil {
		response.BusinessError(&a.Base.Controller, "查询数据出错")
		return
	}

	m := map[string]interface{}{
		"iTotal":        total,
		"iTotalRecords": len(articleList),
		"aData":         articleList,
	}

	response.Success(&a.Base.Controller, m)
}
