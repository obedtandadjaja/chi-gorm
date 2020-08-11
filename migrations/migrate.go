package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/obedtandadjaja/chi-gorm/models"
)

func InitMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
