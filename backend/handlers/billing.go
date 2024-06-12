package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Billing struct {
	gorm.Model
	CustomerID string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
}

func BillingManagement(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			var billings []Billing
			db.Find(&billings)
			c.JSON(http.StatusOK, gin.H{"success": true, "billings": billings})
		} else if c.Request.Method == http.MethodPost {
			var billing Billing
			if err := c.ShouldBindJSON(&billing); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid data"})
				return
			}
			db.Create(&billing)
			c.JSON(http.StatusOK, gin.H{"success": true, "billing": billing})
		}
	}
}
