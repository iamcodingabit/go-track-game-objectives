package handler

import (
	"encoding/json"
	"iamcodingabit/go-track-game-objectives/internal/store"
	"log"
	"net/http"
	"strconv"
)

type GameHandler struct {
	store *store.GameStore
}

func NewGameHandler(s *store.Store) *GameHandler {
	return &GameHandler{
		store: s.Games,
	}
}

func(gh *GameHandler) GetAllGamesHandler(w http.ResponseWriter, r *http.Request){
	games, err := gh.store.GetAllGames(r.Context())
	if err != nil {
		log.Println("fetch games error:", err)
		http.Error(w, "failed to retrieve games", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func (gh *GameHandler) DeleteGameHandler(w http.ResponseWriter, r *http.Request){
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return
	}

	deletedTitle, err := gh.store.DeleteGameById(r.Context(), id)
	if err != nil {
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "deleted successfully",
		"title":   deletedTitle,
	})
}

func(gh *GameHandler) AddGameHandler(w http.ResponseWriter, r *http.Request){
	var newGame store.Game

	err := json.NewDecoder(r.Body).Decode(&newGame)
	if err != nil {
		log.Println("decode error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := gh.store.AddNewGame(r.Context(), newGame)
	if err != nil {
		log.Println("skiwhoa")
		log.Println("decode error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "deleted successfully",
		"title": res,
	})
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