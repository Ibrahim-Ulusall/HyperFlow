package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type BaseRepository[TEntity any] struct {
	Database *gorm.DB
}

func (baseRepository *BaseRepository[TEntity]) GetList() ([]TEntity, error) {
	if baseRepository.Database == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	dbInstance, _ := baseRepository.Database.DB()

	defer func() {
		if err := dbInstance.Close(); err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}()

	var entities []TEntity
	result := baseRepository.Database.Find(&entities)

	if result.Error != nil {
		return nil, fmt.Errorf("error fetching entities: %v", result.Error)
	}

	return entities, nil
}
