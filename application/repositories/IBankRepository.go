package repositories

import (
	core "HyperFlow/core/repositories"
	domain "HyperFlow/domain/entities"
)

type IBankRepository interface {
	core.IRepository[domain.Bank]
}
