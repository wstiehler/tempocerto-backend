package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeHealthHandle(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	})

}
