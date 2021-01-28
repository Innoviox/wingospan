package main

func (g *Game) clone() *Game {
	other := new(Game)


	return other
}

func (p *Player) clone() *Player {
	other := new(Player)
	other.board = p.board.clone()

	other.food = map[Food]int {}
	for i := 0; i < 5; i++ {
		other.food[Food(i)] = p.food[Food(i)]
	}

	other.hand = make([]Bird, len(p.hand))
	for i := 0; i < len(p.hand); i++ {
		other.hand[i] = p.hand[i].clone()
	}

	// todo bonus

	return other
}

func (b *Board) clone() *Board {

}

func (b Bird) clone() Bird {

}

func (g *Game) load(other *Game) {

}