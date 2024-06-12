package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func UserManagement(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			var users []User
			db.Find(&users)
			c.JSON(http.StatusOK, gin.H{"success": true, "users": users})
		} else if c.Request.Method == http.MethodPost {
			var user User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid data"})
				return
			}
			db.Create(&user)
			c.JSON(http.StatusOK, gin.H{"success": true, "user": user})
		}
	}
}
