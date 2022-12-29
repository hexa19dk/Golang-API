package config

import (
	"go-api/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Function to connect to database
func DatabaseConnection() *gorm.DB {
	dsn := "root:DarkRoast123#@tcp(127.0.0.1:3306)/go-api?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&entity.Book{}, &entity.User{})
	return db
}

// Close database connection
func CloseDBConnect(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close database connection")
	}
	dbSQL.Close()
}
