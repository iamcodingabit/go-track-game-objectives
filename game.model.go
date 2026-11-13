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