package database

import (
	"fmt"
	"gofiber-jwt-auth/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)

	var err error
	for retries := 5; retries > 0; retries-- {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Connection Opened to Database")
			break
		}
		log.Printf("Retrying database connection, attempts remaining: %d", retries-1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&model.Book{}, &model.User{})
	if err != nil {
		panic("failed to migrate database schemas")
	}
	fmt.Println("Database Migrated")
}
