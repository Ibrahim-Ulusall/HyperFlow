package entities

import (
	core "HyperFlow/core/entities"
	"time"

	"github.com/google/uuid"
)

type User struct {
	core.BaseEntity[uuid.UUID]
	Firstname                     string     `gorm:"column:firstname"`
	Lastname                      string     `gorm:"column:lastname"`
	Email                         string     `gorm:"column:email"`
	Username                      string     `gorm:"column:Username"`
	Photo                         string     `gorm:"column:photo"`
	Status                        bool       `gorm:"column:status"`
	DepartmentId                  int        `gorm:"column:birim_id"`
	OfisId                        int        `gorm:"column:ofis_id"`
	CompanyId                     int        `gorm:"column:company_id"`
	PasswordHash                  []byte     `gorm:"column:password_hash"`
	PasswordSalt                  []byte     `gorm:"column:password_salt"`
	Token                         string     `gorm:"column:token"`
	TokenExpiration               time.Time  `gorm:"column:token_expiration"`
	ForgotPasswordToken           string     `gorm:"column:forgot_password_token"`
	ForgotPasswordTokenExpiration time.Time  `gorm:"column:forgot_password_token_expiration"`
	TwoFactorEnabled              bool       `gorm:"column:two_factor_enabled"`
	Department                    Department `gorm:"foreignkey:DepartmentId;references:ID"`
	Company                       Company    `gorm:"foreignkey:CompanyId;references:ID"`
	Office                        Office     `gorm:"foreignkey:OfisId;references:ID"`
}

func (User) TableName() string {
	return "accounts.users"
}
