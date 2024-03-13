package tempocerto

import "time"

type ModelsEntity struct {
	CompanyEntity *CompanyEntity
	AvailableSlot *AvailableSlot
}

type CompanyEntity struct {
	ID        uint      `gorm:"primary_key"`
	CNPJ      string    `json:"cnpj" gorm:"unique"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AvailableSlot struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CompanyID uint      `json:"company_id" gorm:"index"`
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	Start     string    `json:"start"`
	End       string    `json:"end"`
	Available string    `json:"available"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WeeklySlotEntity struct {
	ID        uint      `gorm:"primary_key"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	Weekdays  []string  `json:"weekdays"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *CompanyEntity) TableName() string {
	return "companies"
}

func (p *AvailableSlot) TableName() string {
	return "available_slots"
}
