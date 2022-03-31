package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null;"`
	Age   int    `json:"age"`
	Lists []List `json:"lists"` // wont apply cascade delete
}
