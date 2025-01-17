package repositories

import (
	core "HyperFlow/core/repositories"
	"HyperFlow/domain/entities"
)

type DepartmentRepository struct {
	core.BaseRepository[entities.Department]
}
