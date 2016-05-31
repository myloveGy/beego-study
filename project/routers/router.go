package routers

import (
	"github.com/astaxie/beego"
	"project/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexHomeController{})
	beego.Router("/ajax", &controllers.IndexHomeController{}, "*:Ajax")
}
