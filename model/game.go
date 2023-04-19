package model

import (
	"encoding/csv"
	"fmt"
	"github.com/innoviox/wingospan/model/util"
	"math/rand"
	"os"
	"time"
)

type Game struct {
	Players []*Player

	deck []Bird
	tray []Bird

	birdfeeder Birdfeeder

	goals [4]Goal

	round int
}

func (g *Game) Init(nplayers int) {
	rand.Seed(time.Now().Unix())

	g.loadDeck()

	g.tray = g.draw(3)

	for i := 0; i < nplayers; i++ {
		g.Players = append(g.Players, &Player{
			p_idx: i,
			board: new(Board),
			food: map[util.Food]int{
				util.Worm: 1, util.Seed: 1, util.Fish: 1, util.Rodent: 1, util.Berry: 1,
			},
			Hand: g.draw(5),
			// todo bonus
		})
	}

	g.birdfeeder = Birdfeeder{
		diceIn:  []Dice{Die(), Die(), Die(), Die(), Die()},
		diceOut: []Dice{},
	}

	g.chooseGoals()

	// todo goals
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

		g.deck = append(g.deck, Bird{
			Name:     line[0],
			region:   util.ParseRegion(line[1]),
			Cost:     readCost(line[2]),
			points:   util.Atoi(line[3]),
			nest:     util.Nest(util.Atoi(line[4])),
			eggLimit: util.Atoi(line[5]),
			wingspan: util.Atoi(line[6]),
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

func (g *Game) drawTray(idxs []int) []Bird {
	var t []Bird
	var out []Bird

	for i, b := range g.tray {
		if util.Contains(idxs, i) {
			out = append(out, b)
		} else {
			t = append(t, b)
		}
	}

	g.tray = t
	return out
}

func (g *Game) getPlayer(i int) *Player {
	for _, p := range g.Players {
		if p.p_idx == i {
			return p
		}
	}

	return nil
}

func (g *Game) start() {
	for _, p := range g.Players {
		possibleMoves := p.generateMoves(g)
		for i, m := range possibleMoves {
			fmt.Fprintf(os.Stdout, "%d %s\n", i, m.String())
		}
	}
}
