package tempocerto

import (
	"errors"
	"strings"
	"time"

	"github.com/wstiehler/tempocerto-backend/internal/infrastructure/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	repo Repository
}

func NewService(repo *Repository) *Service {
	return &Service{*repo}
}

func NormalizeString(s string) string {
	return strings.ToLower(s)
}

var (
	ErrFailedToCreateCompany = errors.New("failed to create company")
	ErrFailedToCreateSlot    = errors.New("failed to create slot")
)

func (s *Service) CreateCompany(db *gorm.DB, company *CompanyEntity) (*CompanyDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	logger.Debug("Starting to create a new company")

	company.Name = NormalizeString(company.Name)

	createdCompany, err := s.repo.CreateCompany(company)
	if err != nil {
		logger.Error("Failed to create company:", zap.String("error", err.Error()))
		return nil, errors.New("failed to create company")
	}

	createdCompanyDTO := &CompanyDTO{
		ID:   createdCompany.ID,
		CNPJ: createdCompany.CNPJ,
		Name: createdCompany.Name,
	}

	logger.Debug("Company created successfully")

	return createdCompanyDTO, nil
}

func (s *Service) CreateDailyAvailableSlots(db *gorm.DB, date time.Time) ([]AvailableSlotDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	logger.Debug("Starting to create daily available slots")

	var createdSlots []AvailableSlotDTO

	workStartTime, _ := time.Parse("15:04", "08:00")
	workEndTime, _ := time.Parse("15:04", "18:00")
	interval := time.Hour

	for slotStart := workStartTime; slotStart.Before(workEndTime); slotStart = slotStart.Add(interval) {
		slotEnd := slotStart.Add(interval)
		newSlot := &AvailableSlot{
			Date:      date.AddDate(0, 0, 1),
			Start:     slotStart.Format("15:04"),
			End:       slotEnd.Format("15:04"),
			Available: "true",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		logger.Debug("Creating slot", zap.Any("slot", newSlot))

		createdSlot, err := s.repo.CreateAvailableSlot(newSlot)
		if err != nil {
			logger.Error("Error creating available slot:", zap.String("error", err.Error()))
			return nil, err
		}

		createdSlotDTO := &AvailableSlotDTO{
			CompanyID: createdSlot.CompanyID,
			Title:     createdSlot.Title,
			Date:      date,
			Start:     createdSlot.Start,
			End:       createdSlot.End,
			Available: createdSlot.Available,
		}

		createdSlots = append(createdSlots, *createdSlotDTO)
	}

	logger.Info("Daily available slots created successfully")

	return createdSlots, nil
}

func (s *Service) FillSlotByDateTime(db *gorm.DB, date time.Time, start, end, title string, companyID uint) (*AvailableSlotDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	logger.Info("Filling slot by date and time")

	availableSlot, err := s.repo.GetAvailableSlotByDateTime(db, 0, date.AddDate(0, 0, 1), start, end)
	if err != nil {
		logger.Error("Error getting available slot by date and time:", zap.String("error", err.Error()))
		return nil, err
	}

	availableSlot.Available = "false"
	availableSlot.Title = title
	availableSlot.CompanyID = companyID

	if err := s.repo.UpdateAvailableSlot(&availableSlot); err != nil {
		logger.Error("Error updating available slot:", zap.String("error", err.Error()))
		return nil, err
	}

	slotDTO := &AvailableSlotDTO{
		CompanyID: availableSlot.CompanyID,
		Title:     availableSlot.Title,
		Date:      availableSlot.Date,
		Start:     availableSlot.Start,
		End:       availableSlot.End,
		Available: availableSlot.Available,
	}

	logger.Info("Slot filled successfully")

	return slotDTO, nil
}

func (s *Service) GetAllAvailableSlots(db *gorm.DB) ([]*AvailableSlotDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	logger.Debug("Retrieving all available slots")

	slots, err := s.repo.GetAllSlots(db)
	if err != nil {
		logger.Error("Error retrieving all available slots:", zap.String("error", err.Error()))
		return nil, err
	}

	var availableSlots []*AvailableSlotDTO
	for _, slot := range slots {
		availableSlots = append(availableSlots, &AvailableSlotDTO{
			CompanyID: slot.CompanyID,
			Title:     slot.Title,
			Date:      slot.Date,
			Start:     slot.Start,
			End:       slot.End,
			Available: slot.Available,
		})
	}

	logger.Debug("All available slots retrieved successfully")

	return availableSlots, nil
}

func (s *Service) GetAllSchedules(db *gorm.DB) ([]*DTOResponse, error) {
	logger, dispose := logger.New()
	defer dispose()

	logger.Debug("Retrieving all schedules")

	slots, err := s.repo.GetAllSchedules(db)
	if err != nil {
		logger.Error("Error retrieving all schedules:", zap.String("error", err.Error()))
		return nil, err
	}

	var scheduleSlots []*DTOResponse
	for _, slot := range slots {
		companyInfo, err := s.GetCompanyInfoByID(slot.CompanyID)
		if err != nil {
			logger.Error("Error getting company info by ID:", zap.String("error", err.Error()))
			return nil, err
		}

		response := &DTOResponse{
			Title:     slot.Title,
			Start:     slot.Start,
			End:       slot.End,
			Date:      slot.Date,
			Available: slot.Available,
			Company: CompanyDTO{
				ID:   slot.CompanyID,
				Name: companyInfo.Name,
				CNPJ: companyInfo.CNPJ,
			},
		}

		scheduleSlots = append(scheduleSlots, response)
	}

	logger.Debug("All schedules retrieved successfully")

	return scheduleSlots, nil
}

func (s *Service) GetCompanyInfoByID(companyID uint) (*CompanyDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	logger.Debug("Retrieving company info by ID")

	company, err := s.repo.GetCompanyByID(companyID)
	if err != nil {
		logger.Error("Error retrieving company info by ID:", zap.String("error", err.Error()))
		return nil, err
	}

	companyDTO := &CompanyDTO{
		CNPJ: company.CNPJ,
		Name: company.Name,
	}

	logger.Debug("Company info retrieved successfully")

	return companyDTO, nil
}
