package web

import (
	"context"
	"iamcodingabit/go-track-game-objectives/internal/store"
	"iamcodingabit/go-track-game-objectives/internal/web/component"
	"log"
	"net/http"
)

type GameWebHandler struct {
	store *store.GameStore
}

func NewGameWebHandler(s *store.Store) *GameWebHandler {
	return &GameWebHandler{
		store: s.Games,
	}
}

func (gwh *GameWebHandler) GetAllGamesHandler(w http.ResponseWriter, r *http.Request){
	gameList, err := gwh.store.GetAllGames(r.Context())
	if err != nil {
		log.Println(err)
	}

	if err := component.HelloWorld(gameList).Render(context.TODO(), w); err != nil {
		log.Println(err)
	}
}
