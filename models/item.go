package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model

	Name        string `gorm:"not_null"`
	Description string
	Price       int    `gorm:"not_null"`
	Currency    string `gorm:"default:'USD'"`
}
