package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model

	Type           string `gorm:"not_null;default:'customer'"`
	Email          string `gorm:"unique_index"`
	HashedPassword string `gorm:"not_null"`
	Phone          string `gorm:"index"`
	FirstName      string
	Lastname       string
	Data           postgres.Jsonb
	DeactivatedAt  time.Time

	Orders []Order
}
