package service

import (
	"github.com/astaxie/beego/logs"
	"moneyBag/models"
	"moneyBag/mysql"
)

//最新市场价格
func ListNewMarket(name string) ([]models.NewMarket, error) {
	list := make([]models.NewMarket, 0)
	err := mysql.Engine.Desc("id").Find(&list)
	return list, err
}

//添加市场价格
func AddNewMarket(newMarket models.NewMarket) (models.NewMarket, error) {
	var oldNewMarket models.NewMarket
	has, err := mysql.Engine.Where("name=?", newMarket.Name).Get(&oldNewMarket)
	if err != nil {
		logs.Error("异常:", err)
		return newMarket, err
	}
	var history models.HistoryMarket
	if has {
		oldNewMarket.UnitPrice = newMarket.UnitPrice
		_, err = mysql.Engine.ID(oldNewMarket.Id).Update(&oldNewMarket)
		history = models.HistoryMarket{NewMarketId: oldNewMarket.Id, UnitPrice: newMarket.UnitPrice}
	} else {
		_, err = mysql.Engine.Insert(&newMarket)
		history = models.HistoryMarket{NewMarketId: newMarket.Id, UnitPrice: newMarket.UnitPrice}
	}
	_, err = mysql.Engine.Insert(&history)
	return newMarket, err
}
