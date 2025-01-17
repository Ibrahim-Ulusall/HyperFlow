package repositories

import (
	core "HyperFlow/core/repositories"
	"HyperFlow/domain/entities"
)

type ProjectRepository struct {
	core.BaseRepository[entities.Project]
}
