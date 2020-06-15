package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"moneyBag/models"
	"moneyBag/mysql"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.TplName = "add.html"
}

func (c *MainController) Post() {
	pwd := c.GetString("pwd")

	user := c.GetString("user")
	amount := c.GetString("amount")
	incomePay, _ := c.GetInt("incomePay")
	source := c.GetString("source")
	use := c.GetString("use")
	//验证
	if pwd != "901224" {
		c.Data["Mark"] = "密钥错误"
		c.TplName = "index.tpl"
		return
	}
	if amount == "0" {
		c.Data["Mark"] = "金额错误!"
		c.TplName = "index.tpl"
		return
	}
	var mark string
	if incomePay == 1 {
		if source == "" {
			c.Data["Mark"] = "请填写收入来源!"
			c.TplName = "index.tpl"
			return
		}
		mark = "收入"
		use = ""
		c.Data["Mark"] = "哈哈哈哈哈哈哈哈,又收入了一笔!!!"
	} else {
		if use == "" {
			c.Data["Mark"] = "请填写支出用途!"
			c.TplName = "index.tpl"
			return
		}
		c.Data["Mark"] = "555555555,又支出了一笔..."
		mark = "支出"
		source = ""
		incomePay = -incomePay
	}
	describe := strings.Trim(c.GetString("describe"), " ")
	isMust, _ := c.GetInt("isMust")
	bag := models.MoneyBag{User: user, Amount: amount, IncomePay: incomePay, Source: source, Use: use, Mark: mark, Describe: describe, IsMust: isMust}
	fmt.Println(bag)
	mysql.Engine.Insert(&bag)
	c.Data["Website"] = "添加成功"
	c.Data["Email"] = "奥利给"
	c.TplName = "index.tpl"
}
