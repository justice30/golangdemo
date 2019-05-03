package main

import (
	_ "golangdemo/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run("127.0.0.1:8888")
}

