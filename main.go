package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"moneyBag/controllers"
	_ "moneyBag/mysql"
	_ "moneyBag/routers"
)

func main() {
	//logs.Async(1e3)
	logs.SetLevel(logs.LevelDebug)
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+beego.AppConfig.String("logPath")+`"}`)
	if beego.BConfig.RunMode != "prod" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
