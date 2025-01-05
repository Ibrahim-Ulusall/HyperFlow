package entities

import (
	core "HyperFlow/core/entities"
)

type ExpenseType struct {
	core.BaseEntity[int]
	Name string
}
