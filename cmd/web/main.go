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

	handler, err := handler.NewHandler(db)
	gameWebHandler := web.NewGameWebHandler(db)

	router := http.NewServeMux();
	router.HandleFunc("GET /api/games", handler.Game.GetAllGamesHandler)
	router.HandleFunc("GET /api/game/{id}", handler.Game.GetGameByIdHandler)
	router.HandleFunc("DELETE /api/game/{id}", handler.Game.DeleteGameHandler)
	router.HandleFunc("POST /api/game", handler.Game.AddGameHandler)
	router.HandleFunc("PATCH /api/game/{id}", handler.Game.UpdateGameFieldHandler)

	router.HandleFunc("GET /api/platforms", handler.Platform.GetAllPlatformHandler)
	router.HandleFunc("POST /api/platform", handler.Platform.AddNewPlatformHandler)
	//DELETE PLATFORM DETAIL
	//UPDATE PLATFORM DETAIL

	//GET ALL PLAYTHROUGHS
	//GET ONE PLAYTHROUGH WITH OBJECTIVES

	router.HandleFunc("GET /", gameWebHandler.GetAllGamesHandler)


	http.ListenAndServe(":8080", router)
}
