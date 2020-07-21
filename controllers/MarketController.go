package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"moneyBag/models"
	"moneyBag/service"
	"moneyBag/utils"
)

type MarketController struct {
	beego.Controller
}

func (c *MarketController) Get() {
	defer c.ServeJSON()
	name := c.GetString("name")
	list, err := service.ListNewMarket(name)
	if err != nil {
		c.Data["json"] = utils.ErrReturn(err.Error())
	} else {
		c.Data["json"] = utils.SuccessReturn(list)
	}
}

func (c *MarketController) Post() {
	defer c.ServeJSON()
	name := c.GetString("name")
	unitPrice := c.GetString("unitPrice")
	logs.Info("添加价格:", name)
	newMarket := models.NewMarket{Name: name, UnitPrice: unitPrice}
	data, err := service.AddNewMarket(newMarket)
	if err != nil {
		c.Data["json"] = utils.ErrReturn(err.Error())
	} else {
		c.Data["json"] = utils.SuccessReturn(data)
	}
}
