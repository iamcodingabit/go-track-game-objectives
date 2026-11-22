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
	rows, err := gs.db.QueryContext(ctx, `SELECT id, title, genre FROM game`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	games := []Game{}
	
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.ID, &game.Title, &game.Genre); err != nil {
			return nil, err
		}
		games = append(games, game)
	}	

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

func (gs *GameStore) AddNewGame(ctx context.Context, newGame Game) (string, error){
	var title string
	err := gs.db.QueryRowContext(ctx, `INSERT INTO game(title, genre) VALUES ($1, $2) RETURNING title`, newGame.Title, newGame.Genre).Scan(&title)
	if err != nil {
		return "", err
	}

	return title, nil
}

func (gs *GameStore) DeleteGameById(ctx context.Context, id int64) (string, error) {
	var title string
	err := gs.db.QueryRowContext(ctx, `DELETE FROM game WHERE id = $1 RETURNING title`, id).Scan(&title)
	if err != nil {
		return "", err
	}

	return title, nil
}