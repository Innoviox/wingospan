package model

//
//func (p *model.Player) choosePregame() {
//
//}
//
//func (p *model.Player) chooseMove(g *model.Game, maxply int) model.Move {
//	m, _ := p.maximax(g, 0, maxply)
//	return m
//}
//
//func (p *model.Player) maximax(g *model.Game, ply int, maxply int) (model.Move, int) {
//	//s := strings.Repeat("\t", ply)
//	bestScore := 0
//	var bestMove Move
//
//	for _, m := range p.generateMoves(g) {
//		state := g.clone()
//		newPlayer := state.getPlayer(p.p_idx)
//
//		//fmt.Println(s, m.t, m.a)
//		m.f(newPlayer, state, m.a)
//
//		var score int
//		if ply < maxply {
//			_, score = newPlayer.maximax(state, ply+1, maxply) // todo imperfect information
//		} else {
//			score = newPlayer.board.rawScore()
//		}
//
//		if score > bestScore {
//			bestScore = score
//			bestMove = m
//		}
//	}
//
//	return bestMove, bestScore
//}
