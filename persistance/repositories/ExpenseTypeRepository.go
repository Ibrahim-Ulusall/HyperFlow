package repositories

import (
	core "HyperFlow/core/repositories"
	"HyperFlow/domain/entities"
)

type ExpenseTypeRepository struct {
	core.BaseRepository[entities.ExpenseType]
}
