package main

type Player struct {
	board *Board
	food []Food
	hand []Bird
	bonus []Bonus
}

func (p *Player) playBird(b Bird, r Region, f []Food, e EggPayment) {
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