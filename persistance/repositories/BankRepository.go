package repositories

import (
	core "HyperFlow/core/repositories"
	entities "HyperFlow/domain/entities"

	"gorm.io/gorm"
)

type BankRepository struct {
	core.BaseRepository[entities.Bank]
}

func (b *BankRepository) NewBankRepository(db *gorm.DB) *BankRepository {
	b.Database = db
	return b
}
