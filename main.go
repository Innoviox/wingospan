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
	pregame(p, game, pre[rand.Intn(len(pre))].a)

	for _, b := range p.hand {
		fmt.Println(b.name, b.cost)
	}
	fmt.Println(p.food)

	for i := 0; i < 5; i++ {
		m := p.chooseMove(game, 3)

		switch m.t {
		case PlayBird:
			fmt.Println(m.t.String(), m.a.b, m.a.r, m.a.f)
		case GainFood:
			fmt.Println(m.t.String(), m.a.f)
		case LayEggs:
			fmt.Println(m.t.String(), m.a.e)
		case DrawCards:
			fmt.Println(m.t.String(), m.a.tray, m.a.ndeck)
		}

		m.f(p, game, m.a)
	}
}