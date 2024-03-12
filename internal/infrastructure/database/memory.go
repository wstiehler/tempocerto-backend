package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectMemoryDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("../../../../database-test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}

func CloseMemoryDb(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
