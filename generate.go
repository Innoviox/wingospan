package main

import (
	"fmt"
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
