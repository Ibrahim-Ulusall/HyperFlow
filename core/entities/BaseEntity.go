package entities

import "time"

type BaseEntity[T any] struct {
	ID          T
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate time.Time
}
