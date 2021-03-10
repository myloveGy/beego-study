package admin

import (
	"time"
)

type Controller struct {
	CommController
}

func (a *Controller) Site() {
	a.Data["image"] = []int{1, 2, 3, 4, 5, 6, 1, 2}
	a.TplName = "admin/site.html"
}

// 图片上传
func (a *Controller) Upload() {
	filename := "avatar/" + time.Unix(time.Now().Unix(), 0).Format("20060102")
	a.BaseUpload(
		"ace_update_time",
		filename,
		[]string{".jpg", ".jpeg", ".png", ".gif"},
		2097152,
		"update_time",
	)
}
