package database

import (
	"fmt"
	"log"
	"os"

	"github.com/TomislavGalic/CRUDAPI/controllers"
	"github.com/TomislavGalic/CRUDAPI/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func ConnectDatabase() {
	godotenv.Load()

	dbURI := os.Getenv("DB_URL")

	controllers.DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

	controllers.DB.AutoMigrate(&models.Vehicle{})
}
