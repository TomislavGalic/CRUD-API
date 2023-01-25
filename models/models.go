package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model

	Make_name  string `json:"make_name"`
	Model_name string `json:"model_name"`
}
