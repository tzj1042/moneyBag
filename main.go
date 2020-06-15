package main

import (
	"github.com/astaxie/beego"
	_ "moneyBag/mysql"
	_ "moneyBag/routers"
)

func main() {
	beego.Run()
}
