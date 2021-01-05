package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	game := new(Game)
	game.init(5)

	//game.players[0].generatePregame()

	for _, b := range game.deck {
		fmt.Println(b.cost, b.cost.options())
	}
}