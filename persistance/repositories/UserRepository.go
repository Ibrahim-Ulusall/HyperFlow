package repositories

import (
	core "HyperFlow/core/repositories"
	"HyperFlow/domain/entities"
)

type UserRepository struct {
	core.BaseRepository[entities.User]
}
