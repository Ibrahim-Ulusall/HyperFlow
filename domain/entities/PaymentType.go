package entities

import (
	core "HyperFlow/core/entities"

	"github.com/google/uuid"
)

type PaymentType struct {
	core.BaseEntity[int]
	Name      string    `gorm:"column:adi"`
	HesapNo   string    `gorm:"column:hesap_no"`
	UserId    uuid.UUID `gorm:"column:user_id"`
	BankId    int       `gorm:"column:banka_id"`
	IsActive  bool      `gorm:"column:aktif_mi"`
	OrderBy   uint      `gorm:"column:order_by"`
	CompanyId uint      `gorm:"column:company_id"`
	User      User      `gorm:"foreignkey:UserId;references:ID"`
	Banka     Bank      `gorm:"foreignkey:BankId;references:ID"`
}

func (PaymentType) TableName() string {
	return "masraflar.odeme_turus"
}
