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

func (p *Player) generateMoves() []Move {
	moves := make([]Move, 0)

	// play birds
	for _, b := range p.hand {
		for _, r := range b.region {
			for _, f := range b.cost.options() {
				moves = append(moves, Move {
					p.playBird,
					funcArgs { b: b, r: r, f: f },
				})
			}
		}
	}

	// gain food

	// lay eggs

	// draw cards

	return moves
}