package handlers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
)

type Customer struct {
    gorm.Model
    Name  string json:"name"
    Email string json:"email"
}

func CustomerManagement(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.Request.Method == http.MethodGet {
            var customers []Customer
            db.Find(&customers)
            c.JSON(http.StatusOK, gin.H{"success": true, "customers": customers})
        } else if c.Request.Method == http.MethodPost {
            var customer Customer
            if err := c.ShouldBindJSON(&customer); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid data"})
                return
            }
            db.Create(&customer)
            c.JSON(http.StatusOK, gin.H{"success": true, "customer": customer})
        }
    }
}