package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TomislavGalic/CRUDAPI/initializers"
	"github.com/TomislavGalic/CRUDAPI/models"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	initializers.LoadEnv()
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/vehicles", GetVehicles).Methods("GET")
	r.HandleFunc("/createvehicles", CreateVehicles).Methods("POST")
	r.HandleFunc("/getvehicle/{id}", GetVehicle).Methods("GET")
	r.HandleFunc("/updatevehicles/{id}", UpdateVehicles).Methods("PUT")
	r.HandleFunc("/deletevehicles/{id}", DeleteVehicles).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func GetVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var vehicle []models.Vehicle
	DB.Find(&vehicle)
	json.NewEncoder(w).Encode(vehicle)

	json.NewEncoder(w).Encode(&vehicle)
}

func GetVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var vehicle models.Vehicle
	DB.First(&vehicle, params["id"])
	json.NewDecoder(r.Body).Decode(&vehicle)
	DB.Save(&vehicle)
	json.NewEncoder(w).Encode(vehicle)
}

func CreateVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var vehicle models.Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)
	DB.Create(&vehicle)
	json.NewEncoder(w).Encode(vehicle)
}

func UpdateVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var vehicle models.Vehicle
	DB.First(&vehicle, params["id"])
	json.NewDecoder(r.Body).Decode(&vehicle)
	DB.Save(&vehicle)
	json.NewEncoder(w).Encode(vehicle)
}

func DeleteVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var vehicle models.Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)
	DB.Delete(&vehicle, params["id"])
	json.NewEncoder(w).Encode("The user is deleted")
}
