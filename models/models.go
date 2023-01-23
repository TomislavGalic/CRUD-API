package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model

	Make_id  string `json:"make_id"`
	Model_id string `json:"model_id"`
}
