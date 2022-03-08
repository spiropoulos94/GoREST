package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemName string `json:"name" gorm: "not null"`
}
