package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = ""
	c.Data["Email"] = ""
	c.TplName = "home.html"
}
