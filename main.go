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
	router.HandleFunc("GET /game", app.getAllGames)
	router.HandleFunc("POST /game", app.addGameHandler)
	http.ListenAndServe(":8080", router)
}

func(app *App) getAllGames(write http.ResponseWriter, read *http.Request){
	rows, err := app.DB.Query(`SELECT title, genre FROM game`)
	if err != nil {
		log.Println("Failed reading from database")
		return
	}
	defer rows.Close()

	var games []Game
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.Title, &game.Genre); err != nil {
			log.Fatal(err)
		}
		games = append(games, game)
	}	

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(games)
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