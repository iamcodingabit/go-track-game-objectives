package main

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func initDB(){
	var err error;
	DB, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	
}