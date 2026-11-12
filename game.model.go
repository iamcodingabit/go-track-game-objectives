package main

import (
	//"errors"
	//"fmt"
	//"time"
	_ "github.com/mattn/go-sqlite3"
)

type SubObjective struct {
	ID     int  `json:"id"`
	Title	string	`json:"title"`
	Description	string	`json:"description"`
	Completed bool	`json:"completed"`
}

type Objective struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
	SubObjectives []SubObjective `json:"subobjectives"`
}

type Playthrough struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Platform string `json:"platform"`
	Objectives []Objective `json:"objectives"`
}

type Game struct {
	ID int `json:"id,omitempty"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	CreatedAt string `json:"created_at,omitempty"`
}

func(g *Game) AddGame() (Game, error) {
	db := ConnectDB()

	query := `INSERT INTO game(title, genre)
		VALUES(?, ?) RETURNING *`

	statement, err := db.Prepare(query)
	if err != nil {
		return Game{}, err
	}

	defer statement.Close()

	var newGame Game
	err = statement.QueryRow(
		g.Title,
		g.Genre,
	).Scan(
		&newGame.ID,
		&newGame.Title,
		&newGame.Genre,
		&newGame.CreatedAt,
	)

	if err != nil {
		return Game{}, err
	}

	return newGame, nil
}