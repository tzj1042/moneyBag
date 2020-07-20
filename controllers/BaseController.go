package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

// @Title option
// @Description option
// @Success 200 {string} "hello world"
// @router / [options]
func (c *BaseController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}
