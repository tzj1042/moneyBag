package routers

import (
	"github.com/astaxie/beego"
	"moneyBag/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/add", &controllers.MainController{})
}
