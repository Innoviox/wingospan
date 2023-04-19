package model

import "github.com/innoviox/wingospan/model/util"

func (g *Game) clone() *Game {
	other := new(Game)

	other.Players = make([]*Player, 0)
	for i := 0; i < len(g.Players); i++ {
		other.Players = append(other.Players, g.Players[i].clone())
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
	other.p_idx = p.p_idx
	other.board = p.board.clone()

	other.food = map[util.Food]int{}
	for i := 0; i < 5; i++ {
		other.food[util.Food(i)] = p.food[util.Food(i)]
	}

	other.Hand = cloneBirds(p.Hand)

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
		b.Name,
		b.region,
		b.Cost,
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
	var upface []util.Food
	for i := 0; i < len(d.upface); i++ {
		upface = append(upface, d.upface[i])
	}

	return Dice{
		upface,
	}
}

func (g *Game) load(other *Game) {
	g.Players = other.Players
	g.deck = other.deck
	g.tray = other.tray
	g.birdfeeder = other.birdfeeder
	g.goals = other.goals
	g.round = other.round
}

// I FUCKING HATE YOU GO. WHAT FUCKING LANGUAGE DOESNT FUCKING HAVE GENERICS. FICK YOU
func cloneBirds(arr []Bird) []Bird {
	deck := make([]Bird, 0)
	for i := 0; i < len(arr); i++ {
		deck = append(deck, arr[i].clone())
	}
	return deck
}

func cloneDice(arr []Dice) []Dice {
	deck := make([]Dice, 0)
	for i := 0; i < len(arr); i++ {
		deck = append(deck, arr[i].clone())
	}
	return deck
}
