package connection

import (
	"HyperFlow/configurations"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
}

func (database *Database) GetDbConnection() (*gorm.DB, error) {
	config := configurations.Config{}
	setting, _ := config.ConfigureApp()
	conn, err := gorm.Open(postgres.Open(setting.ConnectionString.DefaultConnection), &gorm.Config{})

	if setting.Development.Log.Debug {
		conn.Logger = logger.Default.LogMode(logger.Info)
	}

	if err != nil {
		return nil, fmt.Errorf("could not connect to the database: %v", err)
	}

	return conn, nil
}
