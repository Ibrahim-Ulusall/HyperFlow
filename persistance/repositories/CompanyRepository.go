package repositories

import (
	core "HyperFlow/core/repositories"
	"HyperFlow/domain/entities"
)

type CompanyRepository struct {
	core.BaseRepository[entities.Company]
}
