package model

import (
	"fmt"
	"github.com/innoviox/wingospan/model/util"
	"strings"
)

type Player struct {
	p_idx int

	board *Board
	food  map[util.Food]int
	Hand  []Bird
	bonus []Bonus

	score int
}

func playBird(p *Player, g *Game, args funcArgs) {
	p.payFood(args.f)
	p.payEggs(&args.e)

	p.discard(&args.b)

	p.board.playBird(args.b, args.r)
}

func gainFood(p *Player, g *Game, args funcArgs) {
	for _, f := range args.f { // todo: if can't gain food, gain random
		// todo reroll?
		// todo birdfeeder !!!
		p.food[f]++
		//p.birdfeeder(g, f)
	}

	if args.discardBird != nil {
		p.discard(args.discardBird)
	}

	p.activate(g, util.Forest)
}

func layEggs(p *Player, g *Game, args funcArgs) {
	p.lay(args.e)

	if args.discardFood != nil {
		f := []util.Food{*args.discardFood}
		p.payFood(f)
	}

	p.activate(g, util.Grasslands)
}

func drawCards(p *Player, g *Game, args funcArgs) {
	p.draw(g.drawTray(args.tray))

	p.draw(g.draw(args.ndeck))

	if args.discardEggs != nil {
		p.payEggs(args.discardEggs)
	}

	p.activate(g, util.Waterlands)
}

func (p *Player) payFood(food []util.Food) {
	for _, f := range food {
		p.food[f]--
	}
}

func (p *Player) canPay(food []util.Food) bool {
	// todo 2 foods = 1 food
	totals := map[util.Food]int{}

	for _, f := range food {
		totals[f]++
	}

	for k, v := range totals {
		if v > p.food[k] {
			return false
		}
	}

	return true
}

func (p *Player) payEggs(e *util.Eggs) {
	for _, loc := range *e {
		p.board.rows[loc[0]][loc[1]].eggs -= loc[2]
	}
}

func (p *Player) lay(e util.Eggs) {
	for _, loc := range e {
		p.board.rows[loc[0]][loc[1]].eggs += loc[2]
	}
}

func (p *Player) discard(b *Bird) {
	var hand []Bird

	for _, c := range p.Hand {
		if c.Name != b.Name { // only works because every card is unique
			hand = append(hand, c)
		}
	}

	p.Hand = hand
}

func (p *Player) birdfeeder(g *Game, f util.Food) {
	for _, d := range g.birdfeeder.diceIn {
		if d.hasFood(f) {
			p.food[f]++
			g.birdfeeder.remove(d)
		}
	}
}

func (p *Player) draw(cards []Bird) {
	p.Hand = append(p.Hand, cards...)
}

func (p *Player) activate(g *Game, r util.Region) {
	for i := p.board.r_idxs[r]; i > 0; i-- {
		p.board.rows[r][i-1].activateBrown(g, p)
	}
}

// Mark: Move generation methods
type Pregame struct {
	birdKeep    []int
	foodDiscard []util.Food
	bonusKeep   int // todo
}

func pregame(p *Player, g *Game, f funcArgs) {
	var hand []Bird
	for _, i := range f.p.birdKeep {
		hand = append(hand, p.Hand[i])
	}
	p.Hand = hand

	for _, i := range f.p.foodDiscard {
		p.food[i]--
	}

	//p.bonus = append([]Bonus{}, p.bonus[f.p.bonusKeep])
}

func (p *Player) String() string {
	var br strings.Builder

	// render board
	//for _, row := range p.board.rows {
	//	fmt.Fprintf(&br, "-\n")
	//fmt.Fprintf(&br, util.DisplayBirdArray(row))
	//fmt.Fprintf(&br, "-\n")
	//}

	// render food
	fmt.Fprintf(&br, util.MapToString(p.food))

	// render hand
	//fmt.Fprintf(&br, util.DisplayBirdArray(p.Hand))

	// todo bonus

	return br.String()
}
