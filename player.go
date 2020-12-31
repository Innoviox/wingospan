package main

type Player struct {
	board *Board
	food []Food
	hand []Bird
	bonus []Bonus
}

func (p *Player) playBird(b Bird, r Region, f []Food, e Eggs) {
	p.payFood(f)
	p.payEggs(e)

	p.discard(b)

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
		p.payFood(discard)
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