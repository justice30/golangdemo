package routers

import (
	"github.com/astaxie/beego"
	"golangdemo/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //beego.Router("/users",&controllers.UserController{})
	//beego.Router("/users/:id",&controllers.UserController{})
	//注解路由
	beego.Include(&controllers.UserController{})
}
