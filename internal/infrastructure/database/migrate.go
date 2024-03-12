package config

import (
	"github.com/wstiehler/tempocerto-backend/internal/domain/tempocerto"
	"gorm.io/gorm"
)

func AutoMigrateTables(db *gorm.DB) error {
	err := db.Table("companies").AutoMigrate(&tempocerto.CompanyEntity{})
	if err != nil {
		return err
	}
	err = db.Table("available_slots").AutoMigrate(&tempocerto.AvailableSlot{})
	if err != nil {
		return err
	}
	return nil
}
