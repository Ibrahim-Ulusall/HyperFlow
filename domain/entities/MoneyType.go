package entities

import (
	core "HyperFlow/core/entities"
)

type MoneyType struct {
	core.BaseEntity[int]
	Name   string
	Symbol string
}
