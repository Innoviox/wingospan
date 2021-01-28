package main

func (g *Game) clone() *Game {
	other := new(Game)

	other.players = make([]*Player, len(g.players))
	for i := 0; i < len(g.players); i++ {
		other.players[i] = g.players[i].clone()
	}

	other.deck = cloneBirds(g.deck)
	other.tray = cloneBirds(g.tray)

	other.birdfeeder = g.birdfeeder.clone()

	other.goals = g.goals
	other.round = g.round

	return other
}

func (p *Player) clone() *Player {
	other := new(Player)
	other.board = p.board.clone()

	other.food = map[Food]int{}
	for i := 0; i < 5; i++ {
		other.food[Food(i)] = p.food[Food(i)]
	}

	other.hand = cloneBirds(p.hand)

	// todo bonus

	return other
}

func (b *Board) clone() *Board {
	other := new(Board)

	var rows [3][]Bird
	for i := 0; i < 3; i++ {
		rows[i] = cloneBirds(b.rows[i])
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
	return Bird{
		b.name,
		b.region,
		b.cost,
		b.points,
		b.nest,
		b.eggLimit,
		b.eggs,
		b.wingspan,
		b.caches,
		b.tucks,
		b.action,
	}
}

func (b Birdfeeder) clone() Birdfeeder {
	return Birdfeeder{
		cloneDice(b.diceIn),
		cloneDice(b.diceOut),
	}
}

func (d Dice) clone() Dice {
	var upface []Food
	for i := 0; i < len(d.upface); i++ {
		upface[i] = d.upface[i]
	}

	return Dice{
		upface,
	}
}

func (g *Game) load(other *Game) {
	g.players = other.players
	g.deck = other.deck
	g.tray = other.tray
	g.birdfeeder = other.birdfeeder
	g.goals = other.goals
	g.round = other.round
}
