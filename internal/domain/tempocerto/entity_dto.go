package tempocerto

import "time"

type CompanyDTO struct {
	ID   uint   `json:"id" `
	CNPJ string `json:"cnpj"`
	Name string `json:"name"`
}

type AvailableSlotDTO struct {
	CompanyID uint      `json:"company_id"`
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	Start     string    `json:"start"`
	End       string    `json:"end"`
	Available string    `json:"available"`
}

type DTOResponse struct {
	Title     string     `json:"title"`
	Start     string     `json:"start"`
	End       string     `json:"end"`
	Date      time.Time  `json:"date"`
	Available string     `json:"available"`
	Company   CompanyDTO `json:"company"`
}
