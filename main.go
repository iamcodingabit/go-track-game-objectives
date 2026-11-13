package main

import (
	"log"
	"net/http"
)

func init(){
	MakeMigrations()
}

func main(){
	http.HandleFunc("POST /game", addGameHandler)
	http.ListenAndServe(":8080", nil)
}
func addGameHandler(write http.ResponseWriter, read *http.Request){
	database, error := initDB()
	if(error != nil){
		log.Fatal(error)
	}

	newGame := Game{
		Title: "Dragon's Dogma 2",
		Genre: "Action RPG",
	}

	_, err := database.Exec(`INSERT INTO game(title, genre) VALUES (?, ?)`, newGame.Title, newGame.Genre)
	if err != nil {
		log.Fatal(err)
	}

}