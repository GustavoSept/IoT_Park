package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sqlx.DB

func ConnectToDatabase() {
	databaseURL := "host=postgres user=root password=root dbname=root port=5432 sslmode=disable"
	var err error
	DB, err = sqlx.Connect("postgres", databaseURL)
	if err != nil {
		log.Panic("Error connecting to the database:", err)
	}
}
