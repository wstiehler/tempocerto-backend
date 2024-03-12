package tempocerto

import (
	"time"

	"gorm.io/gorm"
)

type MemorySqlAdapter struct {
}

func (MemorySqlAdapter) CreateCompany(db *gorm.DB, company *CompanyEntity) (CompanyEntity, error) {
	if err := db.Create(company).Error; err != nil {
		return CompanyEntity{}, err
	}
	return *company, nil
}

func (MemorySqlAdapter) CreateAvailableSlot(db *gorm.DB, slot *AvailableSlot) (AvailableSlot, error) {
	if err := db.Create(slot).Error; err != nil {
		return AvailableSlot{}, err
	}
	return *slot, nil
}

func (MemorySqlAdapter) GetAvailableSlotByDateTime(db *gorm.DB, companyID uint, date time.Time, start, end string) (*AvailableSlot, error) {
	var availableSlot AvailableSlot
	if err := db.Where("company_id = ? AND date = ? AND inicio = ? AND fim = ?", companyID, date, start, end).First(&availableSlot).Error; err != nil {
		return nil, err
	}
	return &availableSlot, nil
}
