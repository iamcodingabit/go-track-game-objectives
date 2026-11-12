package main

import (
	"database/sql"
	"log"
)

var db *sql.DB

func initDB(){
	var err error;
	db, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}	
}