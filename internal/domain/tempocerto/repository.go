package tempocerto

import (
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db           *gorm.DB
	queryAdapter QueryAdapter
}

func NewRepository(db *gorm.DB, queryAdapter QueryAdapter) *Repository {
	return &Repository{
		db:           db,
		queryAdapter: queryAdapter,
	}
}

func (r *Repository) CreateCompany(company *CompanyEntity) (CompanyEntity, error) {
	return r.queryAdapter.CreateCompany(r.db, company)
}

func (r *Repository) CreateAvailableSlot(slot *AvailableSlot) (AvailableSlot, error) {
	return r.queryAdapter.CreateAvailableSlot(r.db, slot)
}

func (r *Repository) GetAvailableSlotByDateTime(db *gorm.DB, num uint, date time.Time, start string, end string) (AvailableSlot, error) {
	return r.queryAdapter.GetAvailableSlotByDateTime(r.db, num, date, start, end)
}

func (r *Repository) UpdateAvailableSlot(slot *AvailableSlot) error {
	return r.queryAdapter.UpdateAvailableSlot(r.db, slot)
}

func (r *Repository) GetAllSlots(db *gorm.DB) ([]AvailableSlot, error) {
	return r.queryAdapter.GetAllSlots(r.db)
}

func (r *Repository) GetAllSchedules(db *gorm.DB) ([]AvailableSlot, error) {
	return r.queryAdapter.GetAllSchedules(r.db)
}

func (r *Repository) GetCompanyByID(id uint) (CompanyEntity, error) {
	return r.queryAdapter.GetCompanyByID(r.db, id)
}
