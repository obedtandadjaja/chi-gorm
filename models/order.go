package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model

	Items          []Item
	DiscountCode   string
	AddressStreet1 string `gorm:"not_null"`
	AddressStreet2 string
	AddressCity    string `gorm:"not_null"`
	AddressState   string `gorm:"not_null"`
	AddressZipCode string `gorm:"not_null"`
	AddressCountry string `gorm:"not_null"`
	DeliveryDate   time.Time
	Status         string `gorm:"not_null;default:'pending'"`
}
