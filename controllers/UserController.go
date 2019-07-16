package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("users", c.Get)
}

// @router /users/:id [get]
func (c *UserController) Get(){
	id := c.GetString("id")
	fmt.Println("=================id:"+id)
}
