package main

import (
	"database/sql"
	"log"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", "./app.sqlite")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err.Error())
	}

	log.Println("Database connection successful")

	return db
}

func MakeMigrations() error {
	db := ConnectDB()

	statement := `CREATE TABLE IF NOT EXISTS game(
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL,
		genre TEXT,
		created_at TIMESTAMP DEFAULT DATETIME
	)

	CREATE TABLE IF NOT EXISTS platform(
		id INTEGER PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT DATETIME
	)

	CREATE TABLE IF NOT EXISTS playthrough(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		status TEXT NOT NULL DEFAULT 'backlog' CHECK(status IN ('backlog', 'in_progress', 'completed')),
		game_id INTEGER NOT NULL REFERENCES game(id),
		platform_id INTEGER NOT NULL REFERENCES platform(id),
		created_at TIMESTAMP DEFAULT DATETIME
	)

	CREATE TABLE IF NOT EXISTS objective(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		type TEXT,
		description TEXT,
		priority INTEGER NOT NULL DEFAULT 1 CHECK (priority BETWEEN 1 AND 5),
		is_completed INTEGER NOT NULL DEFAULT 0 CHECK (is_completed IN (0, 1)),
		parent_id INTEGER references objective(id) ON DELETE CASCADE,
		playthrough_id INTEGER NOT NULL REFERENCES playthrough(id),
		created_at TIMESTAMP DEFAULT DATETIME
	)`

	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	return nil
}
