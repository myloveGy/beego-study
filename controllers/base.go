package controllers

import (
	"github.com/astaxie/beego"
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
type BaseController struct {
	beego.Controller
	User User
}

// 验证用户是否已经登录
func (this *BaseController) IsLogin(str string) bool {
	user := this.GetSession(str)
	if user == nil {
		return false
	}

	this.User = user.(User)
	return this.User.UserId > 0
}

func (b *BaseController) Error(code int, msg string, data interface{}) {
	b.Data["json"] = &Response{Code: code, Msg: msg, Data: data}
	b.ServeJSON()
}

func (b *BaseController) Success(data interface{}, message string) {
	if message == "" {
		message = "success"
	}

	b.Data["json"] = &Response{Code: CodeSuccess, Msg: message, Data: data}
	b.ServeJSON()
}
