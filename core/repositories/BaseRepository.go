package repositories

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)

type BaseRepository[TEntity any] struct {
	Database *gorm.DB
}

func (baseRepository *BaseRepository[TEntity]) Add(entity *TEntity) (*TEntity, error) {
	if baseRepository.Database == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	db, _ := baseRepository.Database.DB()

	defer CloseDatabase(db)

	result := baseRepository.Database.Create(&entity)
	if result.Error != nil {
		return nil, fmt.Errorf("hata: %s", result.Error.Error())
	}

	return entity, nil
}

func (baseRepository *BaseRepository[TEntity]) Update(entity TEntity) (*TEntity, error) {
	if baseRepository.Database == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	db, _ := baseRepository.Database.DB()

	defer CloseDatabase(db)

	result := baseRepository.Database.Save(&entity)

	if result != nil {
		return nil, fmt.Errorf("hata : %v", result.Error)
	}

	return &entity, nil

}

func (baseRepository *BaseRepository[TEntity]) Get(predicate func(*gorm.DB) *gorm.DB) (*TEntity, error) {
	if baseRepository.Database == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	db, _ := baseRepository.Database.DB()

	defer CloseDatabase(db)

	query := baseRepository.Database.Model(new(TEntity))

	if predicate != nil {
		query = predicate(query)
	}

	var entity TEntity
	result := query.First(entity)

	if result != nil {
		return nil, fmt.Errorf("failed to fetch entities: %v", result.Error)
	}
	return &entity, nil
}

func (baseRepository *BaseRepository[TEntity]) GetList() ([]TEntity, error) {
	if baseRepository.Database == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	db, _ := baseRepository.Database.DB()

	defer CloseDatabase(db)

	var entities []TEntity
	result := baseRepository.Database.Find(&entities)

	if result.Error != nil {
		return nil, fmt.Errorf("error fetching entities: %v", result.Error)
	}

	return entities, nil
}

func CloseDatabase(db *sql.DB) {

	if err := db.Close(); err != nil {
		fmt.Println("Error closing database connection:", err)
	}
}
