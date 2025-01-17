package repositories

import (
	core "HyperFlow/core/repositories"
	"HyperFlow/domain/entities"
)

type MoneyTypeRepository struct {
	core.BaseRepository[entities.MoneyType]
}
