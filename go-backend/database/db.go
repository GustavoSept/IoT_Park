package database

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectToDatabase() {
	var err error
	databaseURL := "host=postgres user=root password=root dbname=root port=5432 sslmode=disable"

	// Retry logic
	for i := 1; i <= 40; i++ {
		DB, err = sqlx.Connect("postgres", databaseURL)
		if err == nil {
			log.Println("Database connection established")
			return
		}

		log.Printf("Attempt %d: Database connection failed, retrying in %d milliseconds...\n", i, i*200)
		time.Sleep(time.Duration(i*200) * time.Millisecond)
	}

	log.Panic("Error connecting to the database:", err)
}
