package repositories

import (
	core "HyperFlow/core/repositories"
	"HyperFlow/domain/entities"
)

type OfficeRepository struct {
	core.BaseRepository[entities.Office]
}
