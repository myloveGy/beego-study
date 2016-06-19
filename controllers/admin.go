package controllers
import "time"

type AdminController struct {
	CommController
}

func (this *AdminController) Site() {
	this.Data["image"] = []int{1, 2, 3, 4, 5, 6, 1, 2}
	this.TplName = "admin/site.html"
}

// 修改编辑
func (this *AdminController) Update() {
	this.E = ArrError{Status: 0, Msg: "参数为空", Data: nil}
	this.AjaxReturn()
}

// 图片上传
func (this *AdminController) Upload() {
	filename := "avatar/" + time.Unix(time.Now().Unix(), 0).Format("20060102")
	this.BaseUpload("ace_update_time", filename, []string{".jpg", ".jpeg", ".png", ".gif"}, 2097152, "update_time")
}
