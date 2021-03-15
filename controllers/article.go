package controllers

import (
	"log"
	"strconv"

	"project/connection"
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

// Detail 文章详情
func (a *Article) Detail() {
	var (
		article, next, prev models.Article
	)

	a.Data["action"] = "article"
	a.TplName = "article/detail.html"
	id := a.Ctx.Input.Param("0")
	if id == "" {
		return
	}

	article.Id, _ = strconv.ParseInt(id, 10, 64)
	if article.Id > 0 {
		if err := connection.DB.Find(&article); err != nil {
			log.Println(err)
		} else {
			article.SeeNum += 1
			connection.DB.Update(&models.Article{Id: article.Id, SeeNum: article.SeeNum})
		}

	}

	// 上一篇
	_ = connection.DB.Builder(&prev).
		Where("status", 1).
		Where("id", "<", article.Id).
		OrderBy("id", "desc").
		Limit(1).
		One()

	// 下一篇
	_ = connection.DB.Builder(&next).
		Where("status", 1).
		Where("id", ">", article.Id).
		OrderBy("id", "asc").
		Limit(1).
		One()

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
		articleList = make([]*models.Article, 0)
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

	// 查询数据总条数
	total, err = connection.DB.Builder(&articleList).
		Where("status", 1).
		Paginate(start/length+1, length)
	if err != nil {
		response.BusinessError(&a.Base.Controller, "查询错误")
		return
	}

	response.Success(&a.Base.Controller, &response.PageData{
		Total:        total,
		TotalRecords: len(articleList),
		Data:         articleList,
	})
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

	articleList := make([]*models.Article, 0)
	// 分页查询
	total, err = connection.DB.Builder(&articleList).
		Where("status", 1).
		Where("img", "!=", "").
		OrderBy("id", "desc").
		Paginate(start/length+1, length)
	// 查询数据总条数
	if err != nil {
		response.BusinessError(&a.Base.Controller, "查询数据为空")
		return
	}

	response.Success(&a.Base.Controller, &response.PageData{
		Total:        total,
		TotalRecords: len(articleList),
		Data:         articleList,
	})
}
