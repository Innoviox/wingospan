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
			food: map[Food]int {
				Worm: 1, Seed: 1, Fish: 1, Rodent: 1, Berry: 1,
			},
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

func (g *Game) drawTray(idx int) Bird {
	var t []Bird

	for i, b := range g.tray {
		if i != idx {
			t = append(t, b)
		}
	}

	b := g.tray[idx]
	g.tray = t
	return b
}