package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Name  string `gorm: "not null"`
	Items []Item `gorm:"many2many:user_items;"`
}
