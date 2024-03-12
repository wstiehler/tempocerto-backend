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

func TestCreateWeeklyAvailableSlots_Success(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)

	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	adapter := &tempocerto.MemorySqlAdapter{}

	service := tempocerto.NewService(tempocerto.NewRepository(db, adapter))

	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 7) // One week from now

	start := "08:00"
	end := "12:00"
	weekdays := []string{"Monday", "Wednesday", "Friday"} // Specify the weekdays

	createdSlots, err := service.CreateWeeklyAvailableSlots(db, startDate, endDate, start, end, weekdays)

	assert.NoError(t, err)
	assert.NotNil(t, createdSlots)
}

func TestCreateWeeklyAvailableSlots_Error(t *testing.T) {
	db, err := config.ConnectMemoryDb()
	assert.NoError(t, err)

	defer func() {
		err := config.CloseMemoryDb(db)
		assert.NoError(t, err)
	}()

	adapter := &tempocerto.MemorySqlAdapter{}

	service := tempocerto.NewService(tempocerto.NewRepository(db, adapter))

	// Define os parâmetros para criar os slots
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 7)
	start := "07:00" // Horário fora do horário de trabalho
	end := "09:00"
	weekdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	// Tenta criar os slots
	_, err = service.CreateWeeklyAvailableSlots(db, startDate, endDate, start, end, weekdays)

	// Verifica se ocorreu um erro e se é o erro esperado
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "slot time is outside of work hours (8:00 - 18:00)")
}

func TestParseWeekday_Success(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Weekday
	}{
		{"Sunday", time.Sunday},
		{"Monday", time.Monday},
		{"Tuesday", time.Tuesday},
		{"Wednesday", time.Wednesday},
		{"Thursday", time.Thursday},
		{"Friday", time.Friday},
		{"Saturday", time.Saturday},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("ParseWeekday_%s", test.input), func(t *testing.T) {
			weekday, err := tempocerto.ParseWeekday(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, weekday)
		})
	}
}

func TestParseWeekday_Error(t *testing.T) {

	invalidWeekday := "InvalidDay"
	weekday, err := tempocerto.ParseWeekday(invalidWeekday)
	fmt.Println(weekday)
	assert.Error(t, err)
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

	// Criar alguns slots disponíveis usando CreateWeeklyAvailableSlots
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 7)                 // Uma semana a partir de hoje
	weekdays := []string{"Monday", "Wednesday", "Friday"} // Slots disponíveis nas segundas, quartas e sextas
	_, err = service.CreateWeeklyAvailableSlots(db, startDate, endDate, "08:00", "09:00", weekdays)
	assert.NoError(t, err)

	// Chamar o método GetAllAvailableSlots
	allSlots, err := service.GetAllAvailableSlots(db)
	assert.NoError(t, err)

	// Verificar se todos os slots disponíveis foram retornados
	assert.NotEmpty(t, allSlots)
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

	// Chamar o método GetAllAvailableSlots quando não há slots disponíveis
	allSlots, err := service.GetAllAvailableSlots(db)
	assert.NoError(t, err)

	// Verificar se a lista de slots retornada está vazia
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

	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 7)
	start := "08:00"
	end := "09:00"
	weekdays := []string{"Monday", "Wednesday", "Friday"}

	_, err = service.CreateWeeklyAvailableSlots(db, startDate, endDate, start, end, weekdays)
	assert.NoError(t, err)

	retrievedSchedules, err := service.GetAllSchedules(db)
	assert.NoError(t, err)

	assert.Empty(t, retrievedSchedules)
}

// func TestFillSlotByDateTime_Success(t *testing.T) {
// 	db, err := config.ConnectMemoryDb()
// 	assert.NoError(t, err)

// 	defer func() {
// 		err := config.CloseMemoryDb(db)
// 		assert.NoError(t, err)
// 	}()

// 	adapter := &tempocerto.MemorySqlAdapter{}
// 	repo := tempocerto.NewRepository(db, adapter)
// 	service := tempocerto.NewService(repo)

// 	clearDatabase(db)

// 	startDate := time.Now().Truncate(24 * time.Hour)
// 	endDate := startDate.AddDate(0, 0, 7)
// 	start := "08:00"
// 	end := "09:00"
// 	companyId := uint(1)
// 	title := "Carregamento no Fort"
// 	weekdays := []string{"Monday", "Wednesday", "Friday"}

// 	_, err = service.CreateWeeklyAvailableSlots(db, startDate, endDate, start, end, weekdays)
// 	assert.NoError(t, err)

// 	slotDateTime := startDate.Add(time.Hour * 12)

// 	_, err = service.FillSlotByDateTime(db, slotDateTime, start, end, title, companyId)
// 	assert.NoError(t, err)

// }

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

	startDate := time.Now().Truncate(24 * time.Hour)
	endDate := startDate.AddDate(0, 0, 7)
	start := "08:00"
	end := "09:00"
	companyId := uint(1)
	title := "Carregamento no Fort"
	weekdays := []string{"Monday", "Wednesday", "Friday"}

	_, err = service.CreateWeeklyAvailableSlots(db, startDate, endDate, start, end, weekdays)
	assert.NoError(t, err)

	slotDateTime := startDate.Add(time.Hour * 12)

	_, err = service.FillSlotByDateTime(db, slotDateTime, start, end, title, companyId)
	assert.Error(t, err)
}
