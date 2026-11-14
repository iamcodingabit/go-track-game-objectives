package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	DB *sql.DB
}

func main(){
	db, err := initDB()
	if(err != nil){
		log.Fatal(err)
	}
	defer db.Close()

	MakeMigrations(db)

	app := &App{DB: db}

	router := http.NewServeMux();
	router.HandleFunc("GET /", app.helloNewWorldHandler)
	router.HandleFunc("POST /game", app.addGameHandler)
	http.ListenAndServe(":8080", router)
}

func(app *App) helloNewWorldHandler(write http.ResponseWriter, read *http.Request){
	fmt.Println("Hello, new world, all the boys and girls, I get a story to tellllll")
}

func(app *App) addGameHandler(write http.ResponseWriter, read *http.Request){
	var newGame Game

	err := json.NewDecoder(read.Body).Decode(&newGame)
	if err != nil {
		log.Println("decode error:", err)
		http.Error(write, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := app.DB.ExecContext(read.Context(), `INSERT INTO game(title, genre) VALUES (?, ?)`, newGame.Title, newGame.Genre)
	if err != nil {
		log.Println(res)
		return
	}
}