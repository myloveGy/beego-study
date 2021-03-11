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
func (c *Admin) Search() {
	var arr []*models.Admin

	// 查询信息
	search := map[string]string{
		"search":  "username__icontains",
		"user_id": "user_id",
		"status":  "status",
		"orderBy": "user_id",
		"Table":   "admin",
	}

	// 返回信息
	c.BaseSearch(&arr, search, nil)
}

// 修改数据
func (c *Admin) Update() {
	c.BaseUpdate(&models.Admin{}, "admin")
}
