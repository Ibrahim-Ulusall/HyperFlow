package entities

import (
	core "HyperFlow/core/entities"
)

type PaymentType struct {
	core.BaseEntity[int]
	Name string
}
