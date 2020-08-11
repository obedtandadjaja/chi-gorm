package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/obedtandadjaja/chi-gorm/migrations"
)

func main() {
	db, err := gorm.Open(
		"postgres",
		"host=localhost port=5432 user=obedtandadjaja dbname=chi_gorm sslmode=disable",
	)
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err.Error())
	}
	defer db.Close()

	migrations.InitMigrations(db)

	r := InitRouter()

	fmt.Println("Now serving server at :8000")
	http.ListenAndServe(":8000", r)
}
