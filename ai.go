package main

import (
	"fmt"
	"strings"
)

func (p *Player) choosePregame() {

}

func (p *Player) chooseMove(g *Game) Move {
	m, _ := p.maximax(g, 0, 1)
	return m
}

func (p *Player) maximax(g *Game, ply int, maxply int) (Move, int) {
	s := strings.Repeat("\t", ply)
	bestScore := 0
	var bestMove Move

	for _, m := range p.generateMoves(g) {
		state := g.clone()
		newPlayer := state.getPlayer(p.p_idx)

		m.f(newPlayer, m.a)

		fmt.Println(s, m.t, m.a)
		fmt.Println(s, "birds", newPlayer.board.r_idxs)
		fmt.Println("afterplay", newPlayer.board.rows)

		var score int
		if ply < maxply {
			_, score = newPlayer.maximax(g, ply + 1, maxply) // todo imperfect information
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