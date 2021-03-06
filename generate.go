package main

import (
	"github.com/mxschmitt/golang-combinations"
	"strconv"
)

type funcArgs struct {
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

	p Pregame
}

type Move struct {
	t MoveType
	f func(*Player, *Game, funcArgs)
	a funcArgs
}

func (p *Player) generatePregame() []Move {
	moves := make([]Move, 0)

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

			moves = append(moves, Move {
				PreGame,
				pregame,
				funcArgs { p: Pregame { birdKeep, foodDiscard, bonusKeep } },
			})
		}
	}

	return moves
}

func (p *Player) generateMoves(g *Game) []Move {
	moves := make([]Move, 0)

	// play birds
	for _, b := range p.hand {
		for _, r := range b.region {
			for _, f := range b.cost.options() {
				if p.canPay(f) {
					moves = append(moves, Move{
						PlayBird,
						playBird,
						funcArgs{ b: b, r: r, f: f },
					})
				}
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
			GainFood,
			gainFood,
			funcArgs { f: food },
		})
	}

	// lay eggs
	nEggs := amounts[len(p.board.rows[1])] + 1
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
		for i, spot := range comb {
			e[i] = [3]int { int(spot[0] - 48), int(spot[1] - 48), 1 } // todo condense?
		}

		moves = append(moves, Move {
			LayEggs,
			layEggs,
			funcArgs { e: e },
		})
	}

	// draw cards
	nCards := amounts[len(p.board.rows[2])]
	for nTray := 1; nTray <= len(g.tray); nTray++ {
		nDeck := nCards - nTray

		if nDeck < 0 { continue }

		strs := []string { "0", "1", "2" }

		for _, comb := range combinations.Combinations(strs[0:len(g.tray)], nTray) {
			tray := make([]int, nTray)
			for i, c := range comb {
				tray[i] = Atoi(c)
			}

			moves = append(moves, Move {
				DrawCards,
				drawCards,
				funcArgs { tray: tray, ndeck: nDeck },
			})
		}
	}
	moves = append(moves, Move {
		DrawCards,
		drawCards,
		funcArgs { tray: []int {}, ndeck: nCards },
	})

	return moves
}