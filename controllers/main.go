package controllers

import (
	"github.com/beego/beego/v2/server/web"
)


type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "example.com"
	c.Data["Email"] = "admin@example.com"
	c.TplName = "index.tpl"
}
