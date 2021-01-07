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

	p := game.players[0]

	pre := p.generatePregame()
	p.pregame(pre[rand.Intn(len(pre))])

	for _, b := range p.hand {
		fmt.Println(b.name, b.cost)
	}

	for _, m := range p.generateMoves(game) {
		if m.t == PlayBird {
			fmt.Println(m.a.b, m.a.r, m.a.f)
			m.f(m.a)
			break
		}
	}

	for _, m := range p.generateMoves(game) {
		if m.t == GainFood {
			fmt.Println(m.a.e)
			m.f(m.a)
			break
		}
	}
}