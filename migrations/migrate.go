package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/obedtandadjaja/chi-gorm/models"
)

func InitMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Order{})
	db.AutoMigrate(models.Item{})
}
