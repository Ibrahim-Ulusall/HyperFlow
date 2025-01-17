package entities

import (
	core "HyperFlow/core/entities"
)

type MoneyType struct {
	core.BaseEntity[int]
	Name   string `gorm:"column:name"`
	Symbol string `gorm:"column:symbol"`
}

func (MoneyType) TableName() string {
	return "masraflar.money_types"
}
