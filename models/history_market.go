package models

import (
	"time"
)

type HistoryMarket struct {
	Id          int       `json:"id" xorm:"not null pk autoincr comment('ID') INT(10)"`
	NewMarketId int       `json:"newMarketId" xorm:"comment('最新的市场价格') INT(10)"`
	UnitPrice   string    `json:"unitPrice" xorm:"comment('价格') DECIMAL(10,2)"`
	CreateTime  time.Time `json:"createTime" xorm:"default 'CURRENT_TIMESTAMP' created comment('创建时间') DATETIME"`
}
