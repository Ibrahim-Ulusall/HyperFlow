package entities

import (
	core "HyperFlow/core/entities"
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	core.BaseEntity[int]
	PaymentTypeId int
	ExpenseTypeId int
	ProjectId     int
	CompanyId     int
	MoneyTypeId   int
	UserId        uuid.UUID
	Confirmation  bool
	Private       bool
	Payback       bool
	Description   string
	ImageUrl      string
	Amount        float32
	ExpenseDate   time.Time
}
