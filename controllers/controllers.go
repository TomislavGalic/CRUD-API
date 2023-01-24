package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/TomislavGalic/CRUDAPI/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var vehicle []models.Vehicle
	DB.Find(&vehicle)
	json.NewEncoder(w).Encode(vehicle)
}

func GetVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var vehicle models.Vehicle
	DB.First(&vehicle, params["id"])
	json.NewDecoder(r.Body).Decode(&vehicle)
	json.NewEncoder(w).Encode(vehicle)
}

func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var vehicle models.Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)
	DB.Create(&vehicle)
	json.NewEncoder(w).Encode(vehicle)
}

func UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var vehicle models.Vehicle
	DB.First(&vehicle, params["id"])
	json.NewDecoder(r.Body).Decode(&vehicle)
	DB.Save(&vehicle)
	json.NewEncoder(w).Encode(vehicle)
}

func DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var vehicle models.Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)
	DB.Delete(&vehicle, params["id"])
	json.NewEncoder(w).Encode("The user is deleted")
}
