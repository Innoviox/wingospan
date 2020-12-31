package main

type Game struct {
	players []*Player

	deck []Bird
	tray []Bird

	birdfeeder Birdfeeder

	goals [4]Goal
}

func (g *Game) init(nplayers int) {
	g.tray = g.draw(3)

	for i := 0; i < nplayers; i++ {
		g.players = append(g.players, &Player {
			board: new(Board),
			food: []Food { Worm, Seed, Fish, Rodent, Berry },
			hand: g.draw(5),
			// todo bonus
		})
	}

	g.birdfeeder = Birdfeeder {
		diceIn: []Dice { Die(), Die(), Die(), Die(), Die() },
		diceOut: []Dice {},
	}

	// todo goals
}

func (g *Game) draw(n int) []Bird {
	b := make([]Bird, 0)
	for i := 0; i < n; i++ {
		b = append(b, g.deck[0])
		g.deck = g.deck[1:]
	}
	return b
}