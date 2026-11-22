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

func (gs *GameStore) GetGameById(ctx context.Context, id int64) (Game, error){
	var game Game
	row := gs.db.QueryRowContext(ctx, `SELECT id, title, genre FROM game WHERE id = $1`, id)
	
	err := row.Scan(&game.ID, &game.Title, &game.Genre)
	if err != nil{
		return game, err
	}

	return game, nil
}

func (gs *GameStore) AddNewGame(ctx context.Context, newGame Game) (string, error){
	var title string
	err := gs.db.QueryRowContext(ctx, `INSERT INTO game(title, genre) VALUES ($1, $2) RETURNING title`, newGame.Title, newGame.Genre).Scan(&title)
	if err != nil {
		return title, err
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

func(gs *GameStore) UpdateGameField(ctx context.Context, game Game, id int64) (Game, error) {
	query := `
		UPDATE game
		SET title = COALESCE($1, title),
			genre = COALESCE($2, genre)
		WHERE id = $3
		RETURNING id, title, genre
	`
	var updated Game
	err := gs.db.QueryRowContext(ctx, query, game.Title, game.Genre, id).Scan(&updated.ID, &updated.Title, &updated.Genre)
	if err != nil {
		return updated, err
	}

	return updated, nil
}