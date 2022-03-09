package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"Name" gorm:"not null;"`
	Age  int    `json:"age"`
	// Lists []List
}
