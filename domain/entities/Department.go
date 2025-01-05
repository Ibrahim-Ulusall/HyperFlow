package entities

import core "HyperFlow/core/entities"

type Department struct {
	core.BaseEntity[int]
	Name string
}
