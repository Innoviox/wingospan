package main

import (
	"github.com/mxschmitt/golang-combinations"
	"strconv"
)

func (p *Player) generatePregame() []Pregame {
	moves := make([]Pregame, 0)

	s := make([]string, 0)
	for i := 0; i < 5; i++ {
		s = append(s, "f" + strconv.Itoa(i), "b" + strconv.Itoa(i))
	}

	for _, comb := range combinations.Combinations(s, 5) {
		for bonusKeep := 0; bonusKeep < 2; bonusKeep++ {
			birdKeep, foodKeep := make([]int, 0), make([]string, 0)

			for _, k := range comb {
				if k[0] == 'f' {
					foodKeep = append(foodKeep, k)
				} else {
					birdKeep = append(birdKeep, Atoi(string(k[1])))
				}
			}

			foodDiscard := make([]Food, 0)
			for i := 0; i < 5; i++ {
				fd := "f" + strconv.Itoa(i)
				found := false
				for _, fk := range foodKeep {
					if fk == fd {
						found  = true
					}
				}

				if !found {
					foodDiscard = append(foodDiscard, Food(i))
				}
			}

			moves = append(moves, Pregame { birdKeep, foodDiscard, bonusKeep })
		}
	}

	return moves
}

type funcArgs struct {
	g *Game

	// play bird
	b Bird
	r Region
	f []Food

	// lay eggs
	e Eggs

	// draw cards
	tray []int
	ndeck int

	discardBird *Bird
	discardFood *Food
	discardEggs *Eggs
}

type Move struct {
	f func(funcArgs)
	a funcArgs
}

func (p *Player) generateMoves(g *Game) []Move {
	moves := make([]Move, 0)

	// play birds
	for _, b := range p.hand {
		for _, r := range b.region {
			for _, f := range b.cost.options() { // todo make sure cost can be paid
				moves = append(moves, Move {
					p.playBird,
					funcArgs { g: g, b: b, r: r, f: f },
				})
			}
		}
	}

	amounts := [6]int { 1, 1, 2, 2, 3, 3 }

	// gain food
	nFood := amounts[len(p.board.rows[0])]
	for _, comb := range combrep(nFood, []string { "0", "1", "2", "3", "4" }) {
		food := make([]Food, nFood)
		for i, c := range comb {
			food[i] = Food(Atoi(c))
		}

		moves = append(moves, Move {
			p.gainFood,
			funcArgs { g: g, f: food },
		})
	}

	// lay eggs
	nEggs := amounts[len(p.board.rows[1])]
	spots := make([]string, 0)
	for x, r := range p.board.rows {
		for y, b := range r {
			for i := 0; i < (b.eggLimit - b.eggs); i++ {
				spots = append(spots, strconv.Itoa(x) + strconv.Itoa(y))
			}
		}
	}

	for _, comb := range combinations.Combinations(spots, nEggs) {
		var e Eggs = make([][3]int, len(comb))
		for _, spot := range comb {
			e = append(e, [3]int { int(spot[0] - 48), int(spot[1] - 48), 1 }) // todo condense?
		}

		moves = append(moves, Move {
			p.layEggs,
			funcArgs { g: g, e: e },
		})
	}

	// draw cards
	nCards := amounts[len(p.board.rows[2])]
	for nTray := 0; nTray < 3; nTray++ {
		nDeck := nCards - nTray

		for _, comb := range combinations.Combinations([]string { "0", "1", "2" }, nTray) {
			tray := make([]int, nTray)
			for _, c := range comb {
				tray = append(tray, Atoi(c))
			}

			moves = append(moves, Move {
				p.drawCards,
				funcArgs { g: g, tray: tray, ndeck: nDeck },
			})
		}
	}

	return moves
}