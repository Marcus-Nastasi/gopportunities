package config

import (
	"fmt"

	"github.com/Marcus-Nastasi/gopportunities/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	pass   = "123"
	dbname = "opportunities"
)

func connect() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, pass, dbname, port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("error connecting to pg: %v", err)
		return nil, err
	}
	logger.Infof("GORM: connected to %s\n", dbname)
	logger.Infof("Migrating %s...", dbname)
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("error migrating: %v", err)
		return nil, err
	}
	logger.Infof("%s migrated successfully", dbname)
	return db, nil
}
