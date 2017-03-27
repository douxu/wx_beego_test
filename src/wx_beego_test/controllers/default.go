package controllers

import (
	"github.com/astaxie/beego"
)

type MainControllers struct {
	beego.Controller
}

func (c *MainControllers) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
