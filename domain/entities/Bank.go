package entities

import core "HyperFlow/core/entities"

type Bank struct {
	core.BaseEntity[int]
	Name  string
	Color string
}
