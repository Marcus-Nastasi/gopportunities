package config

import "gorm.io/gorm"

var (
	db *gorm.DB
	logger *Logger
)

func Init() (gorm.DB, error) {
	var err error
	// initializing DB
	db, err = connect()
	if err != nil {
		logger.Errorf("Error while initializing database: %v", err)
		return *db, err
	}
	return *db, nil
}

func GetDb() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	return NewLogger(p)
}
