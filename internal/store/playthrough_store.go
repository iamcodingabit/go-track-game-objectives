package store

import (
	"context"
	"database/sql"
)

/*
type Objective struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
	SubObjectives []Objective `json:"subobjectives"`
}

type Playthrough struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Status string `json:"status"`
	Game Game `json:"game"`
	Platform Platform `json:"platform"`
	Objectives []Objective `json:"objectives"`
}
*/

type Playthrough struct {
	ID int64
	Name string
	Description string
	Status string
	GameID int64
	PlatformID int64
}

type PlaythroughStore struct { 
	db *sql.DB
}


func (ps *PlaythroughStore) List(ctx context.Context) ([]Playthrough, error){
	rows, err := ps.db.QueryContext(ctx, `SELECT id, name, description, status, game_id, platform_id FROM playthrough`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	playthroughs := []Playthrough{}
	
	for rows.Next() {
		var playthrough Playthrough
		if err := rows.Scan(&playthrough.ID, &playthrough.Name, &playthrough.Description, &playthrough.Status, &playthrough.GameID, &playthrough.PlatformID); err != nil {
			return nil, err
		}
		playthroughs = append(playthroughs, playthrough)
	}	

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return playthroughs, nil
}