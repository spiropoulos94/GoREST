package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name   string `json:"name" gorm: "not null"`
	ListID uint
}
