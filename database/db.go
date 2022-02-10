package database

import (
	"log"

	"github.com/gesse-carlos/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {
	url := "host=localhost user=root password=password dbname=students_db port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(url))
	if err != nil {
		log.Panic("Connection error")
	}
	DB.AutoMigrate(&models.Student{})
}