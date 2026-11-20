package main

import (
	"log"
	"net/http"
	"iamcodingabit/go-track-game-objectives/internal/handler"
	"iamcodingabit/go-track-game-objectives/internal/store"
)


func main(){
	db, err := store.NewStore("journal.db")
	if(err != nil){
		log.Fatal(err)
	}
	defer db.Close()
	db.MakeMigrations()

	//gameStore := store.NewGameStore()
	gameHandler := handler.NewGameHandler(db)

	router := http.NewServeMux();
	router.HandleFunc("GET /game", gameHandler.GetAllGamesHandler)
	//router.HandleFunc("POST /game", gameHandler.AddGameHandler)
	http.ListenAndServe(":8080", router)
}
