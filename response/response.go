package response

// Response ajax 响应结构
type Response struct {
	Code int         `json:"code"`           // 响应状态
	Msg  string      `json:"msg"`            // 响应消息
	Data interface{} `json:"data,omitempty"` // 响应数据
}

// 用户数据
type User struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
}
