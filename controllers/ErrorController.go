package controllers

import (
	"fire-control/utils"
	"github.com/astaxie/beego"
)

/**
  该控制器处理页面错误请求
*/
type ErrorController struct {
	beego.Controller
}

func (w *ErrorController) Error401() {
	defer w.ServeJSON()
	w.Data["json"] = utils.ErrReturn("未经授权，请求要求验证身份")
}
func (w *ErrorController) Error403() {
	defer w.ServeJSON()
	w.Data["json"] = utils.ErrReturn("服务器拒绝请求")
}
func (w *ErrorController) Error404() {
	defer w.ServeJSON()
	w.Data["json"] = utils.ErrReturn("很抱歉您访问的地址或者方法不存在")

}

func (w *ErrorController) Error500() {
	defer w.ServeJSON()
	w.Data["json"] = utils.ErrReturn("server error")
}

func (w *ErrorController) Error503() {
	defer w.ServeJSON()
	w.Data["json"] = utils.ErrReturn("服务器目前无法使用（由于超载或停机维护）")
}
