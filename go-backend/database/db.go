package database

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectToDatabase() {
	var err error
	databaseURL := "host=postgres user=root password=root dbname=root port=5432 sslmode=disable"

	// Read the INIT_DB environment variable
	initDB := os.Getenv("INIT_DB")

	// Retry logic for database connection
	for i := 1; i <= 40; i++ {
		DB, err = sqlx.Connect("postgres", databaseURL)
		if err == nil {
			log.Println("Database connection established")

			// If INIT_DB is set to "true", run the initialization scripts
			if initDB == "true" {
				err = runDBInitializationScripts(DB)
				if err != nil {
					log.Panic("Error running initialization scripts:", err)
				}
				log.Println("Database initialization scripts executed successfully")
			}

			return
		}

		log.Printf("Attempt %d: Database connection failed, retrying in %d milliseconds...\n", i, i*200)
		time.Sleep(time.Duration(i*200) * time.Millisecond)
	}

	log.Panic("Error connecting to the database:", err)
}

func runDBInitializationScripts(db *sqlx.DB) error {
	path_initScript := filepath.Join("go-backend", "migration", "docker-database-initial.sql")
	path_insertScript := filepath.Join("go-backend", "migration", "docker-database-inserts.sql")

	if err := executeSQLFile(db, path_initScript); err != nil {
		return err
	}
	if err := executeSQLFile(db, path_insertScript); err != nil {
		return err
	}

	return nil
}

func executeSQLFile(db *sqlx.DB, filePath string) error {
	log.Printf("Executing SQL script: %s\n", filePath)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Splitting file into queries
	queries := strings.Split(string(content), ";")

	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}

		if _, err := db.Exec(query); err != nil {
			return err
		}
	}

	return nil
}
