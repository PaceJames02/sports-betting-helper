package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func InitDb() *sql.DB {
	// Get the connection string from environment variables
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		connStr = "postgres://user:password@localhost:5432/my_database?sslmode=disable"
	}

	var db *sql.DB
	var err error

	// Retry connecting to the database
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		log.Printf("Could not connect to database, retrying in 30 seconds... (%d/10)", i+1)
		time.Sleep(30 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database after 10 attempts: %v", err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return db
}
