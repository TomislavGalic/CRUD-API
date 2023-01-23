package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
type Model struct {
	Model_id              int    `json:"id"`
	Model_name            string `json:"model_name"`
	First_production_year int    `json:"first_production_year"`
}

type Make struct {
	Make_id        int    `json:"id"`
	Make_name string `json:"make_name"`
	Country   string `json:"country"`
}

type Inventory struct {
	Inventory_id         int     `json:"id"`
	Vehicle_id int     `json:"vehicle_id"`
	Color_id   int     `json:"color_id"`
	Price      float64 `json:"price"`
}

type Color struct {
	Color_id   int    `json:"id"`
	Name string `json:"name"`
	Code int    `json:"code"`
}
*/

var DB *gorm.DB
var err error

func main() {

	dbURI := "host=localhost user=postgres password=tomis dbname=vehicles port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

	DB.AutoMigrate(&Vehicle{})

	r := mux.NewRouter()
	r.HandleFunc("/vehicles", getVehicles).Methods("GET")
	r.HandleFunc("/createvehicles", createVehicles).Methods("POST")
	r.HandleFunc("/getvehicle/{id}", getVehicle).Methods("GET")
	r.HandleFunc("/updatevehicles/{id}", updateVehicles).Methods("PUT")
	r.HandleFunc("/deletevehicles/{id}", deleteVehicles).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var vehicle []Vehicle
	DB.Find(&vehicle)
	json.NewEncoder(w).Encode(vehicle)

	json.NewEncoder(w).Encode(&vehicle)
}

func getVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var vehicle Vehicle
	DB.First(&vehicle, params["id"])
	json.NewDecoder(r.Body).Decode(&vehicle)
	DB.Save(&vehicle)
	json.NewEncoder(w).Encode(vehicle)
}

func createVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var vehicle Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)
	DB.Create(&vehicle)
	json.NewEncoder(w).Encode(vehicle)
}

func updateVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var vehicle Vehicle
	DB.First(&vehicle, params["id"])
	json.NewDecoder(r.Body).Decode(&vehicle)
	DB.Save(&vehicle)
	json.NewEncoder(w).Encode(vehicle)
}

func deleteVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var vehicle Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)
	DB.Delete(&vehicle, params["id"])
	json.NewEncoder(w).Encode("The user is deleted")
}
