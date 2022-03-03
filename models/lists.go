package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Name  string
	Items []Item `gorm:"many2many:user_items;"`
}
