package entities

import core "HyperFlow/core/entities"

type Company struct {
	core.BaseEntity[int]
	Name        string `gorm:"column:name"`
	Adres       string `gorm:"column:adres"`
	PhoneNumber string `gorm:"column:phone_number"`
}

func (Company) TableName() string {
	return "public.companies"
}
