package routers

import (
	"github.com/astaxie/beego"
	"golangdemo/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/users",&controllers.UserController{})
}
