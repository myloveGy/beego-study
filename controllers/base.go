package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"project/models"
)

const (
	CodeSuccess       = 10000 // 成功状态码
	CodeMissingParams = 40000 // 缺少参数
	CodeInvalidParams = 40001 // 参数错误
	CodeNotLogin      = 40003 // 未登录
	CodeBusinessError = 40004 // 业务错误
	CodeSystemError   = 40005 // 系统错误
)

// 用户数据
type User struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
}

type Response struct {
	Code int         `json:"code"` // 响应状态
	Msg  string      `json:"msg"`  // 响应消息
	Data interface{} `json:"data"` // 响应数据
}

// 继承基础控制器
type Base struct {
	beego.Controller
	User User
}

// 验证用户是否已经登录
func (b *Base) IsLogin(str string) bool {
	user := b.GetSession(str)
	if user == nil {
		return false
	}

	b.User = user.(User)
	return b.User.UserId > 0
}

func (b *Base) Error(code int, msg string, data interface{}) {
	b.Data["json"] = &Response{Code: code, Msg: msg, Data: data}
	b.ServeJSON()
}

func (b *Base) Success(data interface{}, message string) {
	if message == "" {
		message = "success"
	}

	b.Data["json"] = &Response{Code: CodeSuccess, Msg: message, Data: data}
	b.ServeJSON()
}

// 首页控制器
type Controller struct {
	Base
}

func (c *Controller) Prepare() {

	if c.IsAjax() {
		return
	}

	// 使用的布局
	c.Layout = "layout/home.html"
	status := map[string]interface{}{"status": 1}

	// 点击量
	sees, _ := models.GetArticle(status, 6, "-see_num")

	// 热门
	hots, _ := models.GetArticle(status, 6, "-created_at")
	status["recommend"] = 1

	// 推荐
	commList, _ := models.GetArticle(status, 6, "-created_at")

	// 图文推荐
	var imgList []*models.Article
	orm.NewOrm().QueryTable(&models.Article{}).
		Filter("status", 1).
		Exclude("img", "").
		OrderBy("created_at").
		Limit(5).
		All(&imgList)

	// 用户是否已经登录
	c.Data["isLogin"] = c.IsLogin("user")
	c.Data["sees"] = sees
	c.Data["hots"] = hots
	c.Data["commList"] = commList
	c.Data["imgList"] = imgList
	c.Data["user"] = c.User.Username
}
