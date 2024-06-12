package main

import (
	"golang/handlers" // Adjust the import path according to your project structure
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&handlers.Billing{}, &handlers.Payroll{}, &handlers.User{})

	r := gin.Default()

	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.StaticFile("/", "./static/customer.html")
	r.StaticFile("/billing", "./static/billing.html")
	r.StaticFile("/payroll", "./static/payroll.html")
	r.StaticFile("/user", "./static/user.html")

	r.GET("/customer-management", handlers.CustomerManagement(db))
	r.POST("/customer-management", handlers.CustomerManagement(db))

	r.GET("/billing-management", handlers.BillingManagement(db))
	r.POST("/billing-management", handlers.BillingManagement(db))

	r.GET("/payroll-management", handlers.PayrollManagement(db))
	r.POST("/payroll-management", handlers.PayrollManagement(db))

	r.GET("/user-management", handlers.UserManagement(db))
	r.POST("/user-management", handlers.UserManagement(db))

	r.Run(":8080")
}
