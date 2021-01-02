package main

import (
	"encoding/csv"
	"math/rand"
	"os"
)

type Game struct {
	players []*Player

	deck []Bird
	tray []Bird

	birdfeeder Birdfeeder

	goals [4]Goal

	round int
}

func (g *Game) init(nplayers int) {
	g.loadDeck()

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

	g.chooseGoals()

	// todo goals
	// todo shuffle deck
}

func (g *Game) loadDeck() {
	g.deck = make([]Bird, 0)

	file, _ := os.Open("parse/birds.csv")
	r := csv.NewReader(file)

	records, _ := r.ReadAll()

	for i, line := range records {
		if i == 0 {
			continue
		}

		g.deck = append(g.deck, Bird {
			name:     line[0],
			region:   parseRegion(line[1]),
			cost:     readCost(line[2]),
			points:   Atoi(line[3]),
			nest:     Nest(Atoi(line[4])),
			eggLimit: Atoi(line[5]),
			wingspan: Atoi(line[6]),
			action:   readAction(line[8]),
		})
	}

	rand.Shuffle(len(g.deck), func(i, j int) {
		temp := g.deck[i]
		g.deck[i] = g.deck[j]
		g.deck[j] = temp
	})
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