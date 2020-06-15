package models

import (
	"time"
)

type MoneyBag struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('ID') INT(10)"`
	User       string    `json:"user" xorm:"comment('用户') CHAR(3)"`
	Amount     string    `json:"amount" xorm:"not null comment('金额') DECIMAL(10,2)"`
	IncomePay  int       `json:"incomePay" xorm:"not null comment('收支类型1收入2支出') TINYINT(2)"`
	Source     string    `json:"source" xorm:"not null comment('来源') VARCHAR(255)"`
	Use        string    `json:"use" xorm:"not null comment('用途') VARCHAR(255)"`
	Mark       string    `json:"mark" xorm:"not null comment('备注') VARCHAR(255)"`
	Describe   string    `json:"describe" xorm:"not null comment('描述') VARCHAR(255)"`
	IsMust     int       `json:"isMust" xorm:"not null default 0 comment('是否必须') TINYINT(2)"`
	CreateTime time.Time `json:"createTime" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('创建时间') DATETIME"`
}
