package main

func (p *Player) choosePregame() {

}

func (p *Player) chooseMove(g *Game) Move {
	m, _ := p.maximax(g, 0, 3)
	return m
}

func (p *Player) maximax(g *Game, ply int, maxply int) (Move, int) {
	bestScore := 0
	var bestMove Move

	state := g.clone()

	for _, m := range p.generateMoves(g) {
		g.doMove(m)

		var score int
		if ply < maxply {
			_, score = p.maximax(g, ply + 1, maxply)
		} else {
			score = p.board.rawScore()
		}

		if score > bestScore {
			bestScore = score
			bestMove = m
		}

		g.load(state)
	}

	return bestMove, bestScore
}