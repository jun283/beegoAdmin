package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Data["Website"] = "bee.Admin"
	c.Data["Email"] = "astaxie@gmail.com"

	tname := c.Ctx.Input.URL()
	c.TplName = tname[1:]
	if tname == "/" {
		c.TplName = "index.html"
	}

}
