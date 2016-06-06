package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io"
	"math/rand"
	"os"
	"path"
	"project/models"
	"strconv"
	"time"
)

// 首页控制器
type IndexHomeController struct {
	HomeController
}

// 定义图片类型
type Image struct {
	Title string
	Desc  string
	Url   string
}

// 首页显示
func (this *IndexHomeController) Get() {
	var imgs []Image
	// 查询轮播图片
	_, err := orm.NewOrm().Raw("SELECT `title`, `desc`, `url` FROM `my_image` WHERE `type` = ? AND `status` = ?", 1, 1).QueryRows(&imgs)
	if err != nil {
		this.Redirect("/", 302)
	}

	this.Data["images"] = imgs
	this.Data["action"] = "index"
	this.TplName = "home/index.html"
}

// 请求获取文章信息
func (this *IndexHomeController) Ajax() {
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
			_, err = o.Raw("SELECT COUNT(*) AS `length` FROM `my_article` WHERE `status` = ? AND `img` != ?", 1, "").Values(&maps)
			if err == nil {
				// 查询文章
				num, err1 := o.Raw("SELECT `id`, `title`, `content`, `img`, `create_time`, `see_num`, `comment_num` FROM `my_article` WHERE `status` = ? AND `img` != ? ORDER BY `id` DESC LIMIT ?, ?", 1, "", iStart, iLength).QueryRows(&articls)
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

// 用户登录
func (this *IndexHomeController) Login() {
	// 初始化返回
	this.E = ArrError{Status: 0, Msg: "参数为空", Data: nil}
	// 获取参数
	username, password := this.GetString("username"), this.GetString("password")
	if username != "" && password != "" {
		// 密码加密
		h := sha1.New()
		io.WriteString(h, password)
		password = fmt.Sprintf("%x", h.Sum(nil))

		var u User
		// 查询数据
		err := orm.NewOrm().Raw("SELECT `id`, `username`, `email`, `status` FROM `my_admin` WHERE `username` = ? AND `password` = ? LIMIT 1", username, password).QueryRow(&u)
		this.E.Msg = "用户不存在或者密码错误"
		if err == nil {
			this.E.Msg = "对不起！你被管理员封好了 ):"
			if u.Status == 1 {
				// 设置session
				this.SetSession("user", u)
				this.U = u
				this.E.Status = 1
				this.E.Msg = "登录成功 ^.^ !"
				this.E.Data = u.Username
			}
		}
	}

	this.AjaxReturn()
}

// 用户退出
func (this *IndexHomeController) Logout() {
	// 初始化返回
	this.E = ArrError{Status: 0, Msg: "抱歉，您还没有登录呢!", Data: nil}
	if this.GetSession("user") != nil {
		this.DelSession("user")
		this.E.Status = 1
		this.E.Msg = "您已经退出登录"
	}
	this.AjaxReturn()
}

// 获取图片信息
func (this *IndexHomeController) Image() {
	var maps []orm.Params
	if _, err := orm.NewOrm().Raw("SELECT `title`, `url` FROM `my_image` WHERE `status` = ?", 1).Values(&maps); err == nil {
		this.Data["images"] = maps
		this.Data["action"] = "image"
		this.TplName = "home/image.html"
	}
}

// 新增文章信息
func (this *IndexHomeController) Insert() {
	this.E = ArrError{Status: 0, Msg: "抱歉，您还没有登录呢!", Data: nil}
	var article models.Article
	if this.isLogin("user") {
		this.E.Msg = "提交数据为空"
		if err := this.ParseForm(&article); err == nil {
			article.CreateId = this.U.Id
			article.UpdateId = this.U.Id
			Id, err := models.AddArticle(&article)
			if err == nil && Id > 0 {
				article.Id = Id
				this.E.Status = 1
				this.E.Msg = "添加成功"
				this.E.Data = article
			} else {
				this.E.Msg = err.Error()
			}

		} else {
			this.E.Msg = err.Error()
		}
	}
	this.AjaxReturn()
}

// 添加文章信息
func (this *IndexHomeController) Add() {
	this.E = ArrError{Status: 0, Msg: "抱歉，您还没有登录呢!", Data: nil}
	if this.isLogin("user") {
		// 接收参数
		title, desc, stype, url := this.GetString("title"), this.GetString("desc"), this.GetString("type"), this.GetString("url")
		this.E.Msg = "图片标题和地址为空"
		if title != "" && url != "" {
			i, err := strconv.Atoi(stype)
			if err == nil {
				t := time.Now().Unix()
				_, err = orm.NewOrm().Raw("INSERT INTO `my_image` (`title`, `desc`, `url`, `type`, `create_time`, `create_id`, `update_time`, `update_id`) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", title, desc, url, i, t, this.U.Id, t, this.U.Id).Exec()
				if err == nil {
					this.E.Msg = "添加成功"
					this.E.Status = 1
				} else {
					this.E.Msg = err.Error()
				}
			} else {
				this.E.Msg = err.Error()
			}
		}
	}
	this.AjaxReturn()
}

// 判断文件是否在一个数组中
func InArray(arr []string, val string) bool {
	for _, value := range arr {
		if value == val {
			return true
		}
	}

	return false
}

// 获取文件大小的接口
type Sizer interface {
	Size() int64
}

//  判断目录是否存在
func isDirExists(path string) bool {
	file, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return file.IsDir()
	}

	panic("no reached")
}

// 判断文件是否存在
func isFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 图片上传
func (this *IndexHomeController) Upload() {
	oldFile := this.GetString("fileurl")
	this.E = ArrError{Status: 0, Msg: "抱歉，您还没有登录呢!", Data: nil}
	if this.isLogin("user") {
		f, h, err := this.GetFile("image")
		defer f.Close()
		if err == nil {
			file := path.Ext(h.Filename)
			this.E.Msg = "上传文件格式不对"
			if InArray([]string{".jpg", ".jpeg", ".png", ".gif"}, file) {
				this.E.Msg = "上传文件不能超过2M"
				if 1024*1024*2 > f.(Sizer).Size() {
					// 处理上传目录
					datePath := time.Now().Format("200601")
					dirName := "./static/uploads/" + datePath

					// 目录不存在创建
					if !isDirExists(dirName) {
						err = os.MkdirAll(dirName, 0777)
					}

					this.E.Msg = "创建目录失败 :( " + dirName

					// 创建目录失败
					if err == nil {
						// 文件最终保存的地址
						fileName := dirName + "/" + strconv.Itoa(rand.Int()) + file
						err = this.SaveToFile("image", fileName)
						if err == nil {
							// 上传成功删除之前的图片
							if oldFile != "" && isFileExists("./"+oldFile) {
								os.Remove("./" + oldFile)
							}

							this.E.Status = 1
							this.E.Msg = "图片上传成功"
							this.E.Data = map[string]string{"image": h.Filename, "fileDir": fileName[1:]}
						} else {
							this.E.Msg = err.Error()
						}
					}

				}
			}
		} else {
			this.E.Msg = err.Error()
		}
	}

	this.AjaxReturn()
}
