package admin

import (
	"time"

	"github.com/jinxing-go/mysql"

	"project/models"
)

type Admin struct {
	Comm
}

func (a *Admin) Site() {
	a.Data["image"] = []int{1, 2, 3, 4, 5, 6, 1, 2}
	a.TplName = "admin/admin/site.html"
}

func (a *Admin) Index() {
	a.TplName = "admin/admin/index.html"
}

// 图片上传
func (a *Admin) Upload() {
	filename := "avatar/" + time.Unix(time.Now().Unix(), 0).Format("20060102")
	a.BaseUpload(
		"ace_update_time",
		filename,
		[]string{".jpg", ".jpeg", ".png", ".gif"},
		2097152,
		"update_time",
	)
}

// 查询数据
func (a *Admin) Search() {
	arr := make([]*models.Admin, 0)

	// 查询信息
	search := map[string]string{
		"username": "like",
		"email":    "like",
		"user_id":  "=",
		"status":   "=",
	}

	// 返回信息
	a.baseSearch(&arr, search)
}

func (a *Admin) Create() {
	a.baseCreate(&models.Admin{
		LastLoginIp:   a.Ctx.Request.RemoteAddr,
		LastLoginTime: mysql.Now(),
	})
}

func (a *Admin) Update() {
	a.baseUpdate(&models.Admin{})
}

func (a *Admin) Delete() {
	a.baseDelete(&models.Admin{})
}
