package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"

	"github.com/wstiehler/tempocerto-backend/internal/domain/tempocerto"
	"github.com/wstiehler/tempocerto-backend/internal/infrastructure/logger/logwrapper"
)

type Input struct {
	Logger logwrapper.LoggerWrapper
}

func Cron(input Input, db *gorm.DB, service tempocerto.Service) {

	logger := input.Logger

	c := cron.New()

	logger.Info("Starting TempoCerto-CronJob")

	_, err := c.AddFunc("0 0 * * 0", func() {
		// startDate := time.Now()

		// _, err := service.CreateDailyAvailableSlots(db, startDate)
		// if err != nil {
		// 	fmt.Println("Error on create slots:", err)
		// } else {
		// 	fmt.Println("Slots create with success!")
		// }
	})

	if err != nil {
		fmt.Println("Error on schedule the cron:", err)
		return
	}

	c.Start()

	select {}
}
