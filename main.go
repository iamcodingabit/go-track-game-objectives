package main

import (
	"log"
	"net/http"
	"iamcodingabit/go-track-game-objectives/internal/handler"
	"iamcodingabit/go-track-game-objectives/internal/store"
)


func main(){
	db, err := initDB()
	if(err != nil){
		log.Fatal(err)
	}
	defer db.Close()
	MakeMigrations(db)
	
	

	gameStore := store.NewGameStore(db)
	gameHandler := handler.NewGameHandler(gameStore)

	router := http.NewServeMux();
	router.HandleFunc("GET /game", gameHandler.GetAllGamesHandler)
	//router.HandleFunc("POST /game", gameHandler.AddGameHandler)
	http.ListenAndServe(":8080", router)
}
