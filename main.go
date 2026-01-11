package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	initDb()
}

func initDb() {
	// Get the connection string from environment variables
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		connStr = "postgres://user:password@localhost:5432/my_database?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert data into a table
	// _, err = db.Exec("CREATE TABLE users (id SERIAL PRIMARY KEY, name VARCHAR(100))")
	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
