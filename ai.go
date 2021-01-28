package main

func (p *Player) choosePregame() {

}

func (p *Player) chooseMove(g *Game) Move {
	m, _ := p.maximax(g, 0, 5)
	return m
}

func (p *Player) maximax(g *Game, ply int, maxply int) (Move, int) {
	//s := strings.Repeat("\t", ply)
	bestScore := 0
	var bestMove Move

	for _, m := range p.generateMoves(g) {
		state := g.clone()
		newPlayer := state.getPlayer(p.p_idx)

		//fmt.Println(s, m.t, m.a)
		m.f(newPlayer, state, m.a)

		var score int
		if ply < maxply {
			_, score = newPlayer.maximax(state, ply + 1, maxply) // todo imperfect information
		} else {
			score = newPlayer.board.rawScore()
		}

		if score > bestScore {
			bestScore = score
			bestMove = m
		}
	}

	return bestMove, bestScore
}