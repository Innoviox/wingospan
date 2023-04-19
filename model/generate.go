package model

import (
	"fmt"
	"github.com/innoviox/wingospan/model/util"
	"github.com/mxschmitt/golang-combinations"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

type funcArgs struct {
	// play bird
	b Bird
	r util.Region
	f []util.Food

	// lay eggs
	e util.Eggs

	// draw cards
	tray  []int
	ndeck int

	discardBird *Bird
	discardFood *util.Food
	discardEggs *util.Eggs

	p Pregame
}

type Move struct {
	t util.MoveType
	f func(*Player, *Game, funcArgs)
	a funcArgs
}

func GetFunctionName(i interface{}) string {
	// https://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (m *Move) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "%s %s(%s)", m.t, GetFunctionName(m.f), m.a)

	return br.String()
}

func (p *Player) generatePregame() []Move {
	moves := make([]Move, 0)

	s := make([]string, 0)
	for i := 0; i < 5; i++ {
		s = append(s, "f"+strconv.Itoa(i), "b"+strconv.Itoa(i))
	}

	for _, comb := range combinations.Combinations(s, 5) {
		for bonusKeep := 0; bonusKeep < 2; bonusKeep++ {
			birdKeep, foodKeep := make([]int, 0), make([]string, 0)

			for _, k := range comb {
				if k[0] == 'f' {
					foodKeep = append(foodKeep, k)
				} else {
					birdKeep = append(birdKeep, util.Atoi(string(k[1])))
				}
			}

			foodDiscard := make([]util.Food, 0)
			for i := 0; i < 5; i++ {
				fd := "f" + strconv.Itoa(i)
				found := false
				for _, fk := range foodKeep {
					if fk == fd {
						found = true
					}
				}

				if !found {
					foodDiscard = append(foodDiscard, util.Food(i))
				}
			}

			moves = append(moves, Move{
				util.PreGame,
				pregame,
				funcArgs{p: Pregame{birdKeep, foodDiscard, bonusKeep}},
			})
		}
	}

	return moves
}

func (p *Player) generateMoves(g *Game) []Move {
	moves := make([]Move, 0)

	// play birds
	for _, b := range p.Hand {
		for _, r := range b.region {
			for _, f := range b.Cost.options() {
				if p.canPay(f) {
					moves = append(moves, Move{
						util.PlayBird,
						playBird,
						funcArgs{b: b, r: r, f: f},
					})
				}
			}
		}
	}

	amounts := [6]int{1, 1, 2, 2, 3, 3}

	// gain food
	// todo birdfeeder
	nFood := amounts[len(p.board.rows[0])]
	for _, comb := range util.Combrep(nFood, []string{"0", "1", "2", "3", "4"}) {
		food := make([]util.Food, nFood)
		for i, c := range comb {
			food[i] = util.Food(util.Atoi(c))
		}

		moves = append(moves, Move{
			util.GainFood,
			gainFood,
			funcArgs{f: food},
		})
	}

	// lay eggs
	nEggs := amounts[len(p.board.rows[1])] + 1
	spots := make([]string, 0)
	for x, r := range p.board.rows {
		for y, b := range r {
			for i := 0; i < (b.eggLimit - b.eggs); i++ {
				spots = append(spots, strconv.Itoa(x)+strconv.Itoa(y))
			}
		}
	}

	for _, comb := range combinations.Combinations(spots, nEggs) {
		var e util.Eggs = make([][3]int, len(comb))
		for i, spot := range comb {
			e[i] = [3]int{int(spot[0] - 48), int(spot[1] - 48), 1} // todo condense?
		}

		moves = append(moves, Move{
			util.LayEggs,
			layEggs,
			funcArgs{e: e},
		})
	}

	// draw cards
	nCards := amounts[len(p.board.rows[2])]
	for nTray := 1; nTray <= len(g.tray); nTray++ {
		nDeck := nCards - nTray

		if nDeck < 0 {
			continue
		}

		strs := []string{"0", "1", "2"}

		for _, comb := range combinations.Combinations(strs[0:len(g.tray)], nTray) {
			tray := make([]int, nTray)
			for i, c := range comb {
				tray[i] = util.Atoi(c)
			}

			moves = append(moves, Move{
				util.DrawCards,
				drawCards,
				funcArgs{tray: tray, ndeck: nDeck},
			})
		}
	}
	moves = append(moves, Move{
		util.DrawCards,
		drawCards,
		funcArgs{tray: []int{}, ndeck: nCards},
	})

	return moves
}
