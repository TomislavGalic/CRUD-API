package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model

	Make_id  string `json:"make_id"`
	Model_id string `json:"model_id"`
}

type Model struct {
	Model_id              int    `json:"id"`
	Model_name            string `json:"model_name"`
	First_production_year int    `json:"first_production_year"`
}

type Make struct {
	Make_id   int    `json:"id"`
	Make_name string `json:"make_name"`
	Country   string `json:"country"`
}

type Inventory struct {
	Inventory_id int     `json:"id"`
	Vehicle_id   int     `json:"vehicle_id"`
	Color_id     int     `json:"color_id"`
	Price        float64 `json:"price"`
}

type Color struct {
	Color_id int    `json:"id"`
	Name     string `json:"name"`
	Code     int    `json:"code"`
}
