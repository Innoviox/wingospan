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
	fmt.Println(strings.Repeat("\t", ply), "maximaxing", p.board.rawScore())
	bestScore := 0
	var bestMove Move

	state := g.clone()

	for _, m := range p.generateMoves(g) {
		fmt.Println(strings.Repeat("\t", ply), m.t, m.a)
		g.doMove(m)

		var score int
		if ply < maxply {
			_, score = p.maximax(g, ply + 1, maxply) // todo imperfect information
		} else {
			score = p.board.rawScore()
		}

		if score > bestScore {
			bestScore = score
			bestMove = m
		}

		fmt.Println(strings.Repeat("\t", ply), "found", p.board.rawScore())

		g.load(state)
	}

	return bestMove, bestScore
}