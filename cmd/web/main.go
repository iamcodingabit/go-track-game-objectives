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

	gameHandler := handler.NewGameHandler(db)
	gameWebHandler := web.NewGameWebHandler(db)

	router := http.NewServeMux();
	router.HandleFunc("GET /api/game", gameHandler.GetAllGamesHandler)
	router.HandleFunc("GET /api/game/{id}", gameHandler.GetGameByIdHandler)
	router.HandleFunc("DELETE /api/game/{id}", gameHandler.DeleteGameHandler)
	router.HandleFunc("POST /api/game", gameHandler.AddGameHandler)
	router.HandleFunc("PATCH /api/game/{id}", gameHandler.UpdateGameFieldHandler)

	router.HandleFunc("GET /", gameWebHandler.GetAllGamesHandler)


	http.ListenAndServe(":8080", router)
}
