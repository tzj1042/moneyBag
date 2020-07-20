package controllers

import (
	"github.com/astaxie/beego"
	"moneyBag/service"
	"moneyBag/utils"
)

type MarketController struct {
	beego.Controller
}

func (c *MarketController) Get() {
	defer c.ServeJSON()
	name:= c.GetString("name")
	list, err := service.ListNewMarket(name)
	if err != nil {
		c.Data["json"] = utils.ErrReturn(err.Error())
	} else {
		c.Data["json"] = utils.SuccessReturn(list)
	}
}
