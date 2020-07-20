package models

import (
	"time"
)

type NewMarket struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name       string    `json:"name" xorm:"not null comment('名称') VARCHAR(255)"`
	UnitPrice  string    `json:"unitPrice" xorm:"not null comment('价格') DECIMAL(10,2)"`
	UpdateTime time.Time `json:"updateTime" xorm:"updated comment('更新时间') DATETIME"`
	CreateTime time.Time `json:"createTime" xorm:"default 'CURRENT_TIMESTAMP' created comment('创建时间') DATETIME"`
}
