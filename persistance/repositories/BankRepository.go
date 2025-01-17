package repositories

import (
	core "HyperFlow/core/repositories"
	entities "HyperFlow/domain/entities"
)

type BankRepository struct {
	core.BaseRepository[entities.Bank]
}
