package entities

import core "HyperFlow/core/entities"

type Department struct {
	core.BaseEntity[int]
	Name string `gorm:"column:adi"`
}

func (Department) TableName() string {
	return "accounts.birims"
}
