package main

import (
	"github.com/astaxie/beego"
	_ "golangdemo/routers"
)

func main() {
	//设置静态目录，下载目录
	beego.SetStaticPath("/down1", "download1")
	beego.SetStaticPath("/down2", "download2")
	//beego.AddTemplateExt("html")
	beego.Run("127.0.0.1:8888")

}

