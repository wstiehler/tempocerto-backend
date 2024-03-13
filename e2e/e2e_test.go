//go:build e2e
// +build e2e

package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wstiehler/tempocerto-backend/e2e/cmd"
	"github.com/wstiehler/tempocerto-backend/internal/domain/tempocerto"
)

var url string

func readEnv() (string, int) {
	url = os.Getenv("APPLICATION_URL")
	timeoutString := os.Getenv("TEST_TIMEOUT")

	timeout, err := strconv.Atoi(timeoutString)
	if err != nil {
		timeout = 3
	}
	return url, timeout
}

func TestApiHealth(t *testing.T) {
	assert := assert.New(t)
	readEnv()

	t.Run("Health status", func(t *testing.T) {
		client := cmd.NewProjectApi(url)
		health, err := client.ApiHealth()

		fmt.Printf("health: %+v\n", health)
		assert.Equal(err, nil)
	})
}

var companyId uint

func TestMethodC(t *testing.T) {
	assert := assert.New(t)
	readEnv()

	companyEntity := tempocerto.CompanyEntity{
		CNPJ: "77.866.284/0001-28",
		Name: "teste-tempo certo",
	}

	startDate, _ := time.Parse("2006-01-02", "2024-03-12")
	endDate, _ := time.Parse("2006-01-02", "2024-03-12")

	weeklyEntity := tempocerto.WeeklySlotEntity{
		StartDate: startDate,
		EndDate:   endDate,
		StartTime: "08:00",
		EndTime:   "18:00",
		Weekdays:  []string{"Monday", "Tuesday"},
	}

	t.Run("CreateCompany_When_return_must_be_success", func(t *testing.T) {
		client := cmd.NewProjectApi(url)
		company, err := client.CreateCompany(companyEntity)

		companyId = company.ID

		assert.Equal(company.Name, "teste-tempo certo")
		assert.Equal(err, nil)
	})

	t.Run("CreateWeeklySlots_When_return_must_be_success", func(t *testing.T) {
		client := cmd.NewProjectApi(url)
		_, err := client.CreateWeeklySlots(weeklyEntity)

		assert.Equal(err, nil)
	})
}
