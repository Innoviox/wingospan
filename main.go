package main

import (
	"fmt"
	"math/rand"
)

func main() {

	game := new(Game)
	game.init(5)

	p := game.players[0]

	pre := p.generatePregame()
	pregame(p, game, pre[rand.Intn(len(pre))].a)

	for _, b := range p.hand {
		fmt.Println(b.name, b.cost)
	}
	fmt.Println(p.food)

	fmt.Println(displayBirdArray(p.hand))

}
