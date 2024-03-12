package tempocerto

import (
	"time"

	"gorm.io/gorm"
)

type QueryAdapter interface {
	CreateCompany(db *gorm.DB, company *CompanyEntity) (CompanyEntity, error)
	CreateAvailableSlot(db *gorm.DB, slot *AvailableSlot) (AvailableSlot, error)
	UpdateAvailableSlot(db *gorm.DB, slot *AvailableSlot) error
	GetAvailableSlotByDateTime(db *gorm.DB, num uint, date time.Time, start string, end string) (AvailableSlot, error)
	GetAllSlots(db *gorm.DB) ([]AvailableSlot, error)
	GetAllSchedules(db *gorm.DB) ([]AvailableSlot, error)
	GetCompanyByID(db *gorm.DB, id uint) (CompanyEntity, error)
}
