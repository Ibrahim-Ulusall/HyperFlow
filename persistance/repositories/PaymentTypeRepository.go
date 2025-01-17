package repositories

import (
	core "HyperFlow/core/repositories"
	"HyperFlow/domain/entities"
)

type PaymentTypeRepository struct {
	core.BaseRepository[entities.PaymentType]
}
