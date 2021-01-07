package main

type Player struct {
	board *Board
	food map[Food]int
	hand []Bird
	bonus []Bonus

	score int
}

func (p *Player) playBird(args funcArgs) {
	p.payFood(args.f)
	p.payEggs(&args.e)

	p.discard(&args.b)

	p.board.playBird(args.b, args.r)
}

func (p *Player) gainFood(args funcArgs) {
	for _, f := range args.f { // todo: if can't gain food, gain random
	 						   // todo reroll?
		p.birdfeeder(args.g, f)
	}

	if args.discardBird != nil {
		p.discard(args.discardBird)
	}

	p.activate(args.g, Forest)
}

func (p *Player) layEggs(args funcArgs) {
	p.lay(args.e)

	if args.discardFood != nil {
		f := []Food { *args.discardFood }
		p.payFood(f)
	}

	p.activate(args.g, Grasslands)
}

func (p *Player) drawCards(args funcArgs) {
	for _, i := range args.tray {
		p.draw([]Bird { args.g.drawTray(i) })
	}
	p.draw(args.g.draw(args.ndeck))

	if args.discardEggs != nil {
		p.payEggs(args.discardEggs)
	}

	p.activate(args.g, Waterlands)
}

func (p *Player) payFood(food []Food) {
	for _, f := range food {
		p.food[f]--
	}
}

func (p *Player) canPay(food []Food) bool {
	for _, f := range food {
		if p.food[f] == 0 {
			return false
		}
	}
	return true
}

func (p *Player) payEggs(e *Eggs) {
	for _, loc := range *e {
		p.board.rows[loc[0]][loc[1]].eggs -= loc[2]
	}
}

func (p *Player) lay(e Eggs) {
	for _, loc := range e {
		p.board.rows[loc[0]][loc[1]].eggs += loc[2]
	}
}

func (p *Player) discard(b *Bird) {
	var hand []Bird

	for _, c := range p.hand {
		if c.name != b.name { // only works because every card is unique
			hand = append(hand, c)
		}
	}

	p.hand = hand
}

func (p *Player) birdfeeder(g *Game, f Food) {
	for _, d := range g.birdfeeder.diceIn {
		if d.hasFood(f) {
			p.food[f]++
			g.birdfeeder.remove(d)
		}
	}
}

func (p *Player) draw(cards []Bird) {
	p.hand = append(p.hand, cards...)
}

func (p *Player) activate(g *Game, r Region) {
	for i := p.board.r_idxs[r]; i >= 0; i-- {
		p.board.rows[r][i].activateBrown(g, p)
	}
}

// Mark: Move generation methods
type Pregame struct {
	birdKeep []int
	foodDiscard []Food
	bonusKeep int // todo
}

func (p *Player) pregame(d Pregame) {
	var hand []Bird
	for _, i := range d.birdKeep {
		hand = append(hand, p.hand[i])
	}
	p.hand = hand

	for _, i := range d.foodDiscard {
		p.food[i]--
	}

	//p.bonus = append([]Bonus{}, p.bonus[d.bonusKeep])
}