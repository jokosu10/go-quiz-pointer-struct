package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

func (game *Gamer) AddGame(name string) {
	game.Games = append(game.Games, name)
}

func main() {
	gamer := Gamer{Name: "Zelda"}

	gamer.AddGame("Fifa 2020")
	gamer.AddGame("Mario")
	gamer.AddGame("Soccer 2020")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
