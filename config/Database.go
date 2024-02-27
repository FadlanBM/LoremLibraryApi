package config

import (
	"fmt"
	"github.com/api/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDatabase() {
	conn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia%vJakarta", ENV.DB_HOST, ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_DATABASE, ENV.DB_PORT, "%2F")

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to open database", err)
	}

	db.AutoMigrate(models.Borrowers{})
	db.AutoMigrate(models.Employees{})
	db.AutoMigrate(models.Books{})
	db.AutoMigrate(models.Categories{})
	db.AutoMigrate(models.Book_Category{})
	db.AutoMigrate(models.Collection{})
	db.AutoMigrate(models.Lending{})
	db.AutoMigrate(models.ListLending{})
	db.AutoMigrate(models.Fine{})
	db.AutoMigrate(models.Reviews{})

	DB = db
	log.Println("Database loaded")

}
