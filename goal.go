package main

type GoalDef func (*Player) int

type Goal struct {
	sides [2]GoalDef
	chosen int
}

func eggsInRegion(region Region) GoalDef {
	return func (p *Player) int {
		total := 0

		for _, bird := range p.board.rows[region] {
			total += bird.eggs
		}

		return total
	}
}

func birdsInRegion(region Region) GoalDef {
	return func (p *Player) int {
		total := 0

		for _ = range p.board.rows[region] {
			total++
		}

		return total
	}
}

func eggsInNest(nest Nest) GoalDef {
	return func (p *Player) int {
		total := 0

		for _, row := range p.board.rows {
			for _, bird := range row {
				if bird.nest == nest {
					total += bird.eggs
				}
			}
		}

		return total
	}
}

func nestsWithEggs(nest Nest) GoalDef {
	return func (p *Player) int {
		total := 0

		for _, row := range p.board.rows {
			for _, bird := range row {
				if bird.nest == nest {
					total++
				}
			}
		}

		return total
	}
}

func totalBirds() GoalDef {
	return func (p *Player) int {
		total := 0

		for _, row := range p.board.rows {
			for _ = range row {
				total++
			}
		}

		return total
	}
}

func sets() GoalDef {
	return func (p *Player) int {
		return 0 // todo
	}
}