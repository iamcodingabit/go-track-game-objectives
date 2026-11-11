package main

import "fmt"

type subObjective struct {
	ID     int  `json:"id"`
	Title	string	`json:"title"`
	Description	string	`json:"description"`
	Completed bool	`json:"completed"`
}

type objective struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
	SubObjectives []subObjective `json:"subobjectives"`
}

type playthrough struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Platform string `json:"platform"`
	Objectives []objective `json:"objectives"`
}

type game struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Playthroughs []playthrough `json:"playhroughs"`
}

func main(){
	var custom_playthroughs = []playthrough{
			{ID: 0, Title: "Normal Mode, DLC"},
			{ID: 1, Title: "Hard Mode, starting over"},
	}
	
	var custom_games = game{
		ID: 0,
		Title: "Dragon's Dogma 2: Dark Arisen",
		Playthroughs: custom_playthroughs,
	}
	fmt.Print(custom_games, "\n")
}