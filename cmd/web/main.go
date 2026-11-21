package main

import (
	
	"iamcodingabit/go-track-game-objectives/internal/handler"
	"iamcodingabit/go-track-game-objectives/internal/store"
	"iamcodingabit/go-track-game-objectives/internal/web"
	
	"log"
	"net/http"
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

	gameWebHandler := web.NewGameWebHandler(db)

	router := http.NewServeMux();
	router.HandleFunc("GET /api/game", gameHandler.GetAllGamesHandler)
	//router.HandleFunc("POST /game", gameHandler.AddGameHandler)

	router.HandleFunc("GET /", gameWebHandler.GetAllGamesHandler)


	http.ListenAndServe(":8080", router)
}
