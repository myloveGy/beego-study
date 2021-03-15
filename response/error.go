package response

import (
	"fmt"

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

var (
	errorMessageList = map[int]string{
		CodeSuccess:       "success",
		CodeMissingParams: "缺少参数",
		CodeInvalidParams: "参数错误",
		CodeNotLogin:      "未登录",
		CodeBusinessError: "业务错误",
		CodeSystemError:   "系统错误",
	}
)

// Error 响应错误信息
func Error(b *beego.Controller, code int, params ...interface{}) {
	resp := &Response{Code: code}
	l := len(params)
	if l >= 1 {
		if l >= 2 {
			resp.Data = params[1]
		}

		// 错误提示信息
		message := params[0]
		if m, ok := message.(error); ok {
			resp.Msg = m.Error()
		} else if m1, ok1 := message.(string); ok1 {
			resp.Msg = m1
		} else {
			resp.Msg = fmt.Sprintf("%v", message)
		}
	}

	// 没有传递错误消息
	if resp.Msg == "" {
		if message, ok := errorMessageList[code]; ok {
			resp.Msg = message
		}
	}

	b.Data["json"] = resp
	b.ServeJSON()
}

// MissingParams 参数错误
func MissingParams(b *beego.Controller, params ...interface{}) {
	Error(b, CodeMissingParams, params...)
}

// InvalidParams 缺少参数
func InvalidParams(b *beego.Controller, params ...interface{}) {
	Error(b, CodeInvalidParams, params...)
}

// NotLogin 缺少参数
func NotLogin(b *beego.Controller, params ...interface{}) {
	Error(b, CodeNotLogin, params...)
}

// BusinessError 业务错误
func BusinessError(b *beego.Controller, params ...interface{}) {
	Error(b, CodeBusinessError, params...)
}

// SystemError 服务错误
func SystemError(b *beego.Controller, params ...interface{}) {
	Error(b, CodeSystemError, params...)
}

// Success 成功响应
func Success(b *beego.Controller, data interface{}, params ...interface{}) {

	message := "success"
	if len(params) >= 1 {
		if v, ok := params[0].(string); ok {
			message = v
		} else {
			message = fmt.Sprintf("%v", params[0])
		}
	}

	b.Data["json"] = &Response{Code: CodeSuccess, Msg: message, Data: data}
	b.ServeJSON()
}
