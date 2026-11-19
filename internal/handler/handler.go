package handler

import (
	"encoding/json"
	"iamcodingabit/go-track-game-objectives/internal/store"
	"log"
	"net/http"
)

type GameHandler struct {
	games *store.GameStore
}

func NewGameHandler(games *store.GameStore) *GameHandler {
	return &GameHandler{
		games: games,
	}
}

func(gh *GameHandler) GetAllGamesHandler(w http.ResponseWriter, r *http.Request){
	games, err := gh.games.GetAllGames(r.Context())
	if err != nil {
		log.Println("fetch games error:", err)
		http.Error(w, "failed to retrieve games", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

/*
func(gh *GameHandler) AddGameHandler(write http.ResponseWriter, read *http.Request){
	var newGame Game

	
	if err := json.NewDecoder(read.Body).Decode(&newGame); err != nil {
		log.Println("decode error:", err)
		http.Error(write, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := app.Games.db.ExecContext(read.Context(), `INSERT INTO game(title, genre) VALUES (?, ?)`, newGame.Title, newGame.Genre)
	if err != nil {
		log.Println(res)
		return
	}
}
*/