package entities

import (
	core "HyperFlow/core/entities"
)

type ExpenseType struct {
	core.BaseEntity[int]
	Name     string `gorm:"column:adi"`
	IsActive bool   `gorm:"column:aktif_mi"`
	OrderBy  uint   `gorm:"column:order_by"`
}

func (ExpenseType) TableName() string {
	return "masraflar.masraf_turus"
}
