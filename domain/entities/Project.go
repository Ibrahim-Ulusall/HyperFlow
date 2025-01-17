package entities

import (
	core "HyperFlow/core/entities"
)

type Project struct {
	core.BaseEntity[int]
	Name     string `gorm:"column:adi"`
	IsActive bool   `gorm:"column:aktif_mi"`
	OrderBy  int    `gorm:"column:order_by"`
}

func (Project) TableName() string {
	return "masraflar.dahili_projes"
}
