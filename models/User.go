package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not nulll" json:"first_name"`
	LastName  string `gorm:"not nulll" json:"last_name"`
	Emmail    string `gorm:"not nulll;unique_index" json:"emmail"`
	Tasks     []Task `json:"tasks"`
}
