package store

import (
	"context"
	"database/sql"
)

type Game struct {
	ID int `json:"id,omitempty"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	CreatedAt string `json:"created_at,omitempty"`
}

type GameStore struct {
	db *sql.DB
}

func (gs *GameStore) GetAllGames(ctx context.Context) ([]Game, error){
	rows, err := gs.db.Query(`SELECT title, genre FROM game`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	games := []Game{}
	
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.Title, &game.Genre); err != nil {
			return nil, err
		}
		games = append(games, game)
	}	

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}