package routes

import (
	"net/http"
	"time"

	"github.com/wstiehler/tempocerto-backend/internal/domain/tempocerto"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func MakeRoleHandlers(r *gin.Engine, service tempocerto.Service, db *gorm.DB) {

	group := r.Group("/v1")
	{
		group.POST("/company", CreateCompany(service, db))
		group.GET("/schedules:availables", GetAllAvailableSlots(service, db))

		group.PATCH("/schedule", UpdateAvailableSlot(service, db))
		group.POST("/weekly-slots", CreateWeeklyAvailableSlots(service, db))
		group.GET("/schedules", GetAllSchedules(service, db))
	}
}

func CreateCompany(service tempocerto.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var company tempocerto.CompanyEntity

		if err := c.BindJSON(&company); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		companyCreated, err := service.CreateCompany(db, &company)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company"})
			return
		}

		c.JSON(http.StatusCreated, companyCreated)
	}
}

func UpdateAvailableSlot(service tempocerto.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Title     string `json:"title"`
			CompanyID uint   `json:"company_id"`
			Date      string `json:"date"`
			Start     string `json:"start"`
			End       string `json:"end"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		date, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
			return
		}

		createdSlot, err := service.FillSlotByDateTime(db, date, req.Start, req.End, req.Title, req.CompanyID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fill slot"})
			return
		}

		c.JSON(http.StatusOK, createdSlot)
	}
}

func CreateWeeklyAvailableSlots(service tempocerto.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			StartDate string   `json:"start_date"`
			EndDate   string   `json:"end_date"`
			StartTime string   `json:"start_time"`
			EndTime   string   `json:"end_time"`
			Weekdays  []string `json:"weekdays"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		startDate, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
			return
		}

		endDate, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
			return
		}

		createdSlots, err := service.CreateWeeklyAvailableSlots(db, startDate, endDate, req.StartTime, req.EndTime, req.Weekdays)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create weekly available slots"})
			return
		}

		c.JSON(http.StatusCreated, createdSlots)
	}
}

func GetAllAvailableSlots(service tempocerto.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		slots, err := service.GetAllAvailableSlots(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list slots"})
			return
		}

		c.JSON(http.StatusOK, slots)
	}
}

func GetAllSchedules(service tempocerto.Service, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		slots, err := service.GetAllSchedules(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list schedules"})
			return
		}

		c.JSON(http.StatusOK, slots)
	}
}
