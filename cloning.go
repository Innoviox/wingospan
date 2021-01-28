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
	other := new(Board)

	var rows [3][]Bird
	for i := 0; i < 3; i++ {
		rows[i] = make([]Bird, len(b.rows[i]))
		for j, bird := range b.rows[i] {
			rows[i][j] = bird.clone()
		}
	}

	var r_idxs [3]int
	for i := 0; i < 3; i++ {
		r_idxs[i] = b.r_idxs[i]
	}

	other.rows = rows
	other.r_idxs = r_idxs

	return other
}

func (b Bird) clone() Bird {

}

func (g *Game) load(other *Game) {

}