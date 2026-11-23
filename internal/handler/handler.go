package handler

import "iamcodingabit/go-track-game-objectives/internal/store"

type Handler struct {
	Game *GameHandler
	Platform *PlatformHandler
}

func NewHandler(s *store.Store) (*Handler, error) {
	return &Handler{
		Game: &GameHandler{store: s.Games},
		Platform: &PlatformHandler{store: s.Platform},
	}, nil
}