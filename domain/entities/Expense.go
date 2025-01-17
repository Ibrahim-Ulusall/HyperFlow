package entities

import (
	core "HyperFlow/core/entities"
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	core.BaseEntity[int]
	PaymentTypeId int       `gorm:"column:odeme_turu_id"`
	ExpenseTypeId int       `gorm:"column:masraf_turu_id"`
	ProjectId     int       `gorm:"column:dahili_proje_id"`
	CompanyId     int       `gorm:"column:company_id"`
	MoneyTypeId   int       `gorm:"column:money_type"`
	UserId        uuid.UUID `gorm:"column:user_id"`
	Confirmation  bool      `gorm:"column:onay"`
	Private       bool      `gorm:"column:private"`
	Payback       bool      `gorm:"column:geri_odeme"`
	Description   string    `gorm:"column:description"`
	ImageUrl      string    `gorm:"column:image_url"`
	Amount        float32   `gorm:"column:tutar"`
	ExpenseDate   time.Time `gorm:"column:masraf_tarihi"`
}

func (Expense) TableName() string {
	return "masraflar.masrafs"
}
