package tempocerto

import (
	"time"

	"gorm.io/gorm"
)

type MysqlAdapter struct {
}

func (MysqlAdapter) CreateCompany(db *gorm.DB, company *CompanyEntity) (CompanyEntity, error) {
	if err := db.Create(company).Error; err != nil {
		return CompanyEntity{}, err
	}
	return *company, nil
}

func (MysqlAdapter) CreateAvailableSlot(db *gorm.DB, slot *AvailableSlot) (AvailableSlot, error) {
	if err := db.Create(slot).Error; err != nil {
		return AvailableSlot{}, err
	}
	return *slot, nil
}

func (MysqlAdapter) UpdateAvailableSlot(db *gorm.DB, slot *AvailableSlot) error {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&AvailableSlot{}).Where("id = ?", slot.ID).Updates(slot).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (MysqlAdapter) GetAvailableSlotByDateTime(db *gorm.DB, num uint, date time.Time, start, end string) (AvailableSlot, error) {
	var availableSlot AvailableSlot
	if err := db.Where("company_id = ? AND date = ? AND start = ? AND end = ?", num, date, start, end).First(&availableSlot).Error; err != nil {
		return AvailableSlot{}, err
	}
	return availableSlot, nil
}

func (MysqlAdapter) GetAllSlots(db *gorm.DB) ([]AvailableSlot, error) {
	var slots []AvailableSlot
	if err := db.Find(&slots).Error; err != nil {
		return nil, err
	}
	return slots, nil
}

func (MysqlAdapter) GetAllSchedules(db *gorm.DB) ([]AvailableSlot, error) {
	var slots []AvailableSlot
	if err := db.Where("available = ?", "false").Find(&slots).Error; err != nil {
		return nil, err
	}
	return slots, nil
}

func (MysqlAdapter) GetCompanyByID(db *gorm.DB, id uint) (CompanyEntity, error) {
	var company CompanyEntity
	db.Where("id = ?", id).First(&company)
	return company, nil
}
