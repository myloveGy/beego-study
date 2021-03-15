package admin

import (
	"time"

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
		"username": "username__icontains",
		"email":    "email__icontains",
		"user_id":  "user_id",
		"status":   "status",
	}

	// 返回信息
	a.baseSearch(&arr, search)
}

// 修改数据
func (a *Admin) Update() {
	a.baseUpdate(&models.Admin{})
}

func (a *Admin) Delete() {
	a.baseDelete(&models.Admin{})
}
