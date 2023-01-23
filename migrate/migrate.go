package main

import (
	"github.com/TomislavGalic/CRUDAPI/initializers"
	"github.com/TomislavGalic/CRUDAPI/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Vehicle{}, models.Inventory{}, models.Make{}, models.Color{}, models.Model{})
}
