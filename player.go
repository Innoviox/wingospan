package main

type Player struct {
	board *Board
	food map[Food]int
	hand []Bird
	bonus []Bonus
}

func (p *Player) playBird(b Bird, r Region, f []Food, e Eggs) {
	p.payFood(f)
	p.payEggs(&e)

	p.discard(&b)

	p.board.playBird(b, r)
}

func (p *Player) gainFood(g *Game, food []Food, discard *Bird) {
	for _, f := range food {
		p.birdfeeder(g, f)
	}

	if discard != nil {
		p.discard(discard)
	}

	p.activate(g, Forest)
}

func (p *Player) layEggs(g *Game, e Eggs, discard *Food) {
	p.lay(e)

	if discard != nil {
		f := []Food { *discard }
		p.payFood(f)
	}

	p.activate(g, Grasslands)
}

func (p *Player) drawCards(g *Game, tray []int, ndeck int, discard *Eggs) {
	for _, i := range tray {
		p.draw(g.drawTray(i))
	}
	p.draw(g.draw(ndeck))

	if discard != nil {
		p.payEggs(discard)
	}

	p.activate(g, Waterlands)
}

func (p *Player) payFood(food []Food) {
	for _, f := range food {
		p.food[f]--
	}
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