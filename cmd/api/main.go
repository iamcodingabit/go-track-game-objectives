package main

import (
	"context"
	"iamcodingabit/go-track-game-objectives/internal/handler"
	"iamcodingabit/go-track-game-objectives/internal/store"
	"iamcodingabit/go-track-game-objectives/internal/web/component"
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

	router := http.NewServeMux();
	router.HandleFunc("GET /api/game", gameHandler.GetAllGamesHandler)
	//router.HandleFunc("POST /game", gameHandler.AddGameHandler)

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request){
		if err := component.HelloWorld(db, r.Context()).Render(context.TODO(), w); err != nil {
			log.Println(err)
		}
	})


	http.ListenAndServe(":8080", router)
}
