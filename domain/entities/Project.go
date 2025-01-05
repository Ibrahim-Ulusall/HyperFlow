package entities

import (
	core "HyperFlow/core/entities"
)

type Project struct {
	core.BaseEntity[int]
	Name     string
	IsActive bool
	OrderBy  int
}
