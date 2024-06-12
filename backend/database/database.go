package database

import (
	"handlers" // Adjust the import path according to your project structure
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := "user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Run migrations
	db.AutoMigrate(&handlers.Billing{}, &handlers.Payroll{}, &handlers.User{})

	// Seed initial data
	SeedInitialData(db)

	return db
}

func SeedInitialData(db *gorm.DB) {
	// Check if users table is empty
	var count int64
	db.Model(&handlers.User{}).Count(&count)
	if count == 0 {
		// Create initial users
		password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		initialUsers := []handlers.User{
			{Username: "admin", Password: string(password), Role: "admin"},
			{Username: "sales", Password: string(password), Role: "sales"},
			{Username: "accountant", Password: string(password), Role: "accountant"},
			{Username: "hr", Password: string(password), Role: "hr"},
		}
		for _, user := range initialUsers {
			db.Create(&user)
		}
	}
}
