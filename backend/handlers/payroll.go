package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Payroll struct {
	gorm.Model
	EmployeeID string  `json:"employee_id"`
	Salary     float64 `json:"salary"`
}

func PayrollManagement(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			var payrolls []Payroll
			db.Find(&payrolls)
			c.JSON(http.StatusOK, gin.H{"success": true, "payrolls": payrolls})
		} else if c.Request.Method == http.MethodPost {
			var payroll Payroll
			if err := c.ShouldBindJSON(&payroll); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid data"})
				return
			}
			db.Create(&payroll)
			c.JSON(http.StatusOK, gin.H{"success": true, "payroll": payroll})
		}
	}
}
