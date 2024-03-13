//go:build unit
// +build unit

package tempocerto_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wstiehler/tempocerto-backend/internal/domain/tempocerto"
	config "github.com/wstiehler/tempocerto-backend/internal/infrastructure/database"
	"gorm.io/gorm"
)

func TestCreateCompany_Success(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)

	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	adapter := &tempocerto.MemorySqlAdapter{}

	service := tempocerto.NewService(tempocerto.NewRepository(db, adapter))

	company := &tempocerto.CompanyEntity{
		Name: "Test Company",
		CNPJ: "12345678901234",
	}

	createdCompany, err := service.CreateCompany(db, company)

	assert.NoError(t, err)

	assert.Equal(t, company.Name, createdCompany.Name)
	assert.Equal(t, company.CNPJ, createdCompany.CNPJ)
}
func TestCreateDailyAvailableSlots_Success(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)

	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	adapter := &tempocerto.MemorySqlAdapter{}
	service := tempocerto.NewService(tempocerto.NewRepository(db, adapter))

	date := time.Date(2024, time.March, 13, 0, 0, 0, 0, time.UTC)

	createdSlots, err := service.CreateDailyAvailableSlots(db, date)

	assert.NoError(t, err)
	assert.NotNil(t, createdSlots)

	expectedSlots := 10 // (18:00 - 08:00) / 1 hour interval
	assert.Equal(t, expectedSlots, len(createdSlots))

	workStartTime, _ := time.Parse("15:04", "08:00")
	for i, slot := range createdSlots {
		expectedStartTime := workStartTime.Add(time.Hour * time.Duration(i))
		expectedEndTime := expectedStartTime.Add(time.Hour)

		assert.Equal(t, date, slot.Date)
		assert.Equal(t, expectedStartTime.Format("15:04"), slot.Start)
		assert.Equal(t, expectedEndTime.Format("15:04"), slot.End)
		assert.Equal(t, "true", slot.Available)
	}
}

func TestGetCompanyInfoByID_Success(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)

	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	adapter := &tempocerto.MemorySqlAdapter{}
	repo := tempocerto.NewRepository(db, adapter)
	service := tempocerto.NewService(repo)

	company := &tempocerto.CompanyEntity{
		Name: "Test Company",
		CNPJ: "12345678901234",
	}
	createdCompany, err := service.CreateCompany(db, company)
	assert.NoError(t, err)

	companyID := createdCompany.ID
	companyDTO, err := service.GetCompanyInfoByID(companyID)
	fmt.Println(companyDTO)
	assert.NoError(t, err)
}

func TestGetAllAvailableSlots_Success(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)

	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	adapter := &tempocerto.MemorySqlAdapter{}
	repo := tempocerto.NewRepository(db, adapter)
	service := tempocerto.NewService(repo)

	date := time.Date(2024, time.March, 12, 0, 10, 0, 0, time.UTC)

	_, err = service.CreateDailyAvailableSlots(db, date)
	assert.NoError(t, err)

	allSlots, err := service.GetAllAvailableSlots(db)
	assert.NoError(t, err)

	expectedSlots := 20
	assert.Equal(t, expectedSlots, len(allSlots))

	for _, slot := range allSlots {
		assert.Equal(t, "true", slot.Available)
	}
}

func clearDatabase(db *gorm.DB) {
	db.Exec("DELETE FROM available_slots")
}

func TestGetAllAvailableSlots_Empty(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)

	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	clearDatabase(db)

	adapter := &tempocerto.MemorySqlAdapter{}
	repo := tempocerto.NewRepository(db, adapter)
	service := tempocerto.NewService(repo)

	allSlots, err := service.GetAllAvailableSlots(db)
	assert.NoError(t, err)

	assert.Empty(t, allSlots)
}

func TestGetAllSchedules_Success(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)

	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	adapter := &tempocerto.MemorySqlAdapter{}
	repo := tempocerto.NewRepository(db, adapter)
	service := tempocerto.NewService(repo)

	clearDatabase(db)

	date := time.Date(2024, time.March, 13, 0, 0, 0, 0, time.UTC)
	_, err = service.CreateDailyAvailableSlots(db, date)
	assert.NoError(t, err)

	retrievedSchedules, err := service.GetAllSchedules(db)
	assert.NoError(t, err)

	assert.Empty(t, retrievedSchedules)
}

func TestFillSlotByDateTime_Failure(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)

	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	adapter := &tempocerto.MemorySqlAdapter{}
	repo := tempocerto.NewRepository(db, adapter)
	service := tempocerto.NewService(repo)

	clearDatabase(db)

	date := time.Date(2024, time.March, 13, 0, 0, 0, 0, time.UTC)
	_, err = service.CreateDailyAvailableSlots(db, date)
	assert.NoError(t, err)

	companyId := uint(1)
	title := "Carregamento no Fort"

	slotDateTime := date.Add(time.Hour * 12)

	_, err = service.FillSlotByDateTime(db, slotDateTime, "08:00", "09:00", title, companyId)
	assert.Error(t, err)
}

func TestFillSlotByDateTime_Success(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)

	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	adapter := &tempocerto.MemorySqlAdapter{}
	repo := tempocerto.NewRepository(db, adapter)
	service := tempocerto.NewService(repo)

	date := time.Date(2024, time.March, 13, 0, 0, 0, 0, time.UTC)

	slotDTO, err := service.FillSlotByDateTime(db, date, "08:00", "09:00", "Test Title", 1)
	assert.NoError(t, err)
	assert.NotNil(t, slotDTO)
}

func TestFillSlotByDateTime_ErrorGetAvailableSlot(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)

	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	adapter := &tempocerto.MemorySqlAdapter{}
	repo := tempocerto.NewRepository(db, adapter)
	service := tempocerto.NewService(repo)

	_, err = service.FillSlotByDateTime(db, time.Now(), "08:00", "09:00", "Test Title", 1)
	assert.Error(t, err)
}
