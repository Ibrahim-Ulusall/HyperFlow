package entities

import "time"

type BaseEntity[T any] struct {
	ID          T         `gorm:"primaryKey;autoIncrement;olumn:id"`
	CreatedDate time.Time `gorm:"column:created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date"`
	DeletedDate time.Time `gorm:"column:deleted_date"`
}
