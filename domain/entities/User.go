package entities

import (
	core "HyperFlow/core/entities"
	"time"

	"github.com/google/uuid"
)

type User struct {
	core.BaseEntity[uuid.UUID]
	Firstname                     string
	Lastname                      string
	Email                         string
	Photo                         string
	Status                        bool
	DepartmentId                  int
	CompanyId                     int
	PasswordHash                  []byte
	PasswordSalt                  []byte
	Token                         string
	TokenExpiration               time.Time
	ForgotPasswordToken           string
	ForgotPasswordTokenExpiration time.Time
	TwoFactorEnabled              bool
}
