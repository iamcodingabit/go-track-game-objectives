package main

import (
	"database/sql"
	"log"
	"os"
	_ "modernc.org/sqlite"
)



func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "journal.db")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err.Error())
	}

	log.Println("Database connection successful")

	return db, err;
}


func MakeMigrations() {
	db, err := initDB()
	schemaBytes, err := os.ReadFile("./db_queries/schema.sql")
	if err != nil {
		log.Fatalf ("Failed to load schema.sql: %v", err)
	}

	statement := string(schemaBytes)

	_, err = db.Exec(statement)
	if err != nil {
		log.Fatalf("SOMETHING'S WRONG, I CAN FEEL IT! %s", err)
	}

	log.Println("Migration successful!")
}
