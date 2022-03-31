package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`
	Items  []Item `json:"items" gorm:"constraint:OnDelete:CASCADE;"`
	UserID uint
}
