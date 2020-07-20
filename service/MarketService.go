package service

import (
	"moneyBag/models"
	"moneyBag/mysql"
)

//最新市场价格
func ListNewMarket(name string) ([]models.NewMarket, error) {
	list := make([]models.NewMarket, 0)
	err := mysql.Engine.Desc("id").Find(&list)
	return list, err
}
