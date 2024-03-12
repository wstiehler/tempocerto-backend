package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"

	"github.com/wstiehler/tempocerto-backend/internal/domain/tempocerto"
)

func Cron(db *gorm.DB, service tempocerto.Service, companyID uint) {

	c := cron.New()

	_, err := c.AddFunc("0 0 * * 0", func() {
		// startDate := time.Now()
		// endDate := startDate.AddDate(0, 0, 7)
		// startTime := "08:00"
		// endTime := "18:00"
		// weekdays := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday}

		// _, err := service.CreateWeeklyAvailableSlots(db, companyID, startDate, endDate, startTime, endTime, weekdays)
		// if err != nil {
		// 	fmt.Println("Erro ao criar slots semanais:", err)
		// } else {
		// 	fmt.Println("Slots semanais criados com sucesso!")
		// }
	})

	if err != nil {
		fmt.Println("Erro ao agendar o cronjob:", err)
		return
	}

	c.Start()

	select {}
}
