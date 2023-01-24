package main

import (
	"log"
	"net/http"

	"github.com/TomislavGalic/CRUDAPI/controllers"
	"github.com/TomislavGalic/CRUDAPI/database"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	database.ConnectDatabase()

	r := mux.NewRouter()
	r.HandleFunc("/getvehicles", controllers.GetVehicles).Methods("GET")
	r.HandleFunc("/createvehicle", controllers.CreateVehicle).Methods("POST")
	r.HandleFunc("/getvehicle/{id}", controllers.GetVehicle).Methods("GET")
	r.HandleFunc("/updatevehicle/{id}", controllers.UpdateVehicle).Methods("PUT")
	r.HandleFunc("/deletevehicle/{id}", controllers.DeleteVehicle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
