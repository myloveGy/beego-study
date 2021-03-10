package controllers

import (
	"crypto/sha1"
	"fmt"
	"io"

	"github.com/astaxie/beego/orm"

	"project/models"
)

// 首页控制器
type IndexController struct {
	HomeController
}

// 定义图片类型
type Image struct {
	Title string
	Desc  string
	Url   string
}

// 首页显示
func (i *IndexController) Get() {
	var imgList []Image

	// 查询轮播图片
	orm.NewOrm().Raw("SELECT `title`, `desc`, `url` FROM `my_image` WHERE `type` = ? AND `status` = ?", 1, 1).QueryRows(&imgList)

	i.Data["images"] = imgList
	i.Data["action"] = "index"
	i.TplName = "home/index.html"
}

// 请求获取文章信息
func (i *IndexController) Article() {
	// 初始化返回
	var iStart, iLength int
	var err error
	var articleList []models.Article
	// 接收参数
	iStart, err = i.GetInt("iStart")
	if err != nil {
		i.Error(CodeMissingParams, "参数为空", nil)
		return
	}

	iLength, err = i.GetInt("iLength")
	if err != nil {
		i.Error(CodeMissingParams, "参数为空", nil)
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
	_, err = o.Raw("SELECT COUNT(*) AS `length` FROM `my_article` WHERE `status` = ? AND `img` != ?", 1, "").Values(&maps)
	if err != nil {
		i.Error(CodeBusinessError, "查询数据为空", nil)
		return
	}

	// 查询文章
	num, err1 := o.Raw(
		"SELECT `id`, `title`, `content`, `img`, `create_time`, `see_num`, `comment_num` FROM `my_article` WHERE `status` = ? AND `img` != ? ORDER BY `id` DESC LIMIT ?, ?",
		1,
		"",
		iStart,
		iLength,
	).QueryRows(&articleList)
	if err1 != nil {
		i.Error(CodeBusinessError, "查询数据出错", nil)
		return
	}

	m["iTotal"] = maps[0]["length"]
	m["iTotalRecords"] = num
	i.Success(m, "")
}

// 用户登录
func (i *IndexController) Login() {
	// 获取参数
	username, password := i.GetString("username"), i.GetString("password")
	if username == "" && password == "" {
		i.Error(CodeMissingParams, "参数为空", nil)
		return
	}

	// 密码加密
	h := sha1.New()
	io.WriteString(h, password)
	password = fmt.Sprintf("%x", h.Sum(nil))

	var u User
	// 查询数据
	if err := orm.NewOrm().Raw(
		"SELECT `id`, `username`, `email`, `status` FROM `my_admin` WHERE `username` = ? AND `password` = ? LIMIT 1",
		username,
		password).QueryRow(&u); err != nil {
		i.Error(CodeInvalidParams, "用户不存在或者密码错误", nil)
		return
	}

	if u.Status != 1 {
		i.Error(CodeBusinessError, "对不起！你被管理员封好了 ):", nil)
		return
	}

	// 设置session
	i.SetSession("user", u)
	i.User = u
	i.Success(map[string]interface{}{
		"username": username,
		"email":    u.Email,
		"user_id":  u.Id,
	}, "登录成功")
}

// 用户退出
func (i *IndexController) Logout() {
	// 初始化返回
	i.DelSession("user")
	i.Success(nil, "您已经退出登录")
}

// 获取图片信息
func (i *IndexController) Image() {
	var maps []orm.Params
	if _, err := orm.NewOrm().Raw("SELECT `title`, `url` FROM `my_image` WHERE `status` = ?", 1).Values(&maps); err == nil {
		i.Data["images"] = maps
	}

	i.Data["action"] = "image"
	i.TplName = "home/image.html"
}
