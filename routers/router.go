package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"moneyBag/controllers"
)

func init() {
	//  用于跨域请求
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		//AllowMethods:     []string{"*"},
		//AllowHeaders:     []string{"*"},
		//ExposeHeaders:    []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.NSRouter("*", &controllers.BaseController{}, "OPTIONS:Options")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/add", &controllers.MainController{})
	beego.Router("/newMarket", &controllers.MarketController{})
}
