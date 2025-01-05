package entities

import core "HyperFlow/core/entities"

type Company struct {
	core.BaseEntity[int]
	Name string
}
