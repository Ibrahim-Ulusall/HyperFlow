package repositories

import (
	core "HyperFlow/core/repositories"
	"HyperFlow/domain/entities"
)

type ExpenseRepository struct {
	core.BaseRepository[entities.Expense]
}
