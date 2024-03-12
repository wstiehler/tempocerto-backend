package tempocerto

import (
	"errors"
	"fmt"
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

func (s *Service) CreateWeeklyAvailableSlots(db *gorm.DB, startDate time.Time, endDate time.Time, start, end string, weekdays []string) ([]*AvailableSlotDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	logger.Debug("Starting to create weekly available slots")

	var createdSlots []*AvailableSlotDTO

	workStartTime, _ := time.Parse("15:04", "08:00")
	workEndTime, _ := time.Parse("15:04", "18:00")
	slotStart, _ := time.Parse("15:04", start)
	slotEnd, _ := time.Parse("15:04", end)
	if slotStart.Before(workStartTime) || slotEnd.After(workEndTime) {
		errMsg := "slot time is outside of work hours (8:00 - 18:00)"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	var parsedWeekdays []time.Weekday
	for _, wdStr := range weekdays {
		weekday, err := parseWeekday(wdStr)
		if err != nil {
			logger.Error("Error parsing weekday:", zap.String("error", err.Error()))
			return nil, err
		}
		parsedWeekdays = append(parsedWeekdays, weekday)
	}

	interval := time.Hour

	for d := startDate; d.Before(endDate) || d.Equal(endDate); d = d.AddDate(0, 0, 1) {
		for _, weekday := range parsedWeekdays {
			if d.Weekday() == weekday {
				for slotStart := workStartTime; slotStart.Before(workEndTime); slotStart = slotStart.Add(interval) {
					slotEnd := slotStart.Add(interval)
					newSlot := &AvailableSlot{
						Date:      d,
						Start:     slotStart.Format("15:04"),
						End:       slotEnd.Format("15:04"),
						Available: "true",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					}

					createdSlot, err := s.repo.CreateAvailableSlot(newSlot)
					if err != nil {
						logger.Error("Error creating available slot:", zap.String("error", err.Error()))
						return nil, err
					}

					createdSlotDTO := &AvailableSlotDTO{
						CompanyID: createdSlot.CompanyID,
						Title:     createdSlot.Title,
						Date:      createdSlot.Date,
						Start:     createdSlot.Start,
						End:       createdSlot.End,
						Available: createdSlot.Available,
					}

					createdSlots = append(createdSlots, createdSlotDTO)
				}
			}
		}
	}

	logger.Info("Weekly available slots created successfully")

	return createdSlots, nil
}

func parseWeekday(weekdayStr string) (time.Weekday, error) {
	switch weekdayStr {
	case "Sunday":
		return time.Sunday, nil
	case "Monday":
		return time.Monday, nil
	case "Tuesday":
		return time.Tuesday, nil
	case "Wednesday":
		return time.Wednesday, nil
	case "Thursday":
		return time.Thursday, nil
	case "Friday":
		return time.Friday, nil
	case "Saturday":
		return time.Saturday, nil
	default:
		return time.Sunday, fmt.Errorf("invalid weekday: %s", weekdayStr)
	}
}

func (s *Service) FillSlotByDateTime(db *gorm.DB, date time.Time, start, end, title string, companyID uint) (*AvailableSlotDTO, error) {
	logger, dispose := logger.New()
	defer dispose()

	logger.Info("Filling slot by date and time")

	availableSlot, err := s.repo.GetAvailableSlotByDateTime(db, 0, date, start, end)
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
