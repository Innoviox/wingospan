package main

type Player struct {
	board *Board
	food []Food
	hand []Bird
	bonus []Bonus
}

func (p *Player) playBird(g *Game, b Bird, r Region, f []Food, e EggPayment) {
	p.payFood(f)
	p.payEggs(e)

	p.discard(b)

	p.board.playBird(b, r)
}