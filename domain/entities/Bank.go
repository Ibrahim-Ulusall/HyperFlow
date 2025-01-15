package entities

import core "HyperFlow/core/entities"

type Bank struct {
	core.BaseEntity[int]
	Name  string `gorm:"column:adi"`
	Color string `gorm:"column:color"`
}

func (Bank) TableName() string {
	return "masraflar.bankas"
}
