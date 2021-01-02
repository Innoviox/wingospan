package main

import (
	"math/rand"
)

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

func allGoals() []Goal {
	goals := make([]Goal, 0)

	goals = append(goals, Goal { [2]GoalDef { eggsInRegion(Forest), birdsInRegion(Forest) }, rand.Intn(2) })
	goals = append(goals, Goal { [2]GoalDef { eggsInRegion(Grasslands), birdsInRegion(Grasslands) }, rand.Intn(2) })
	goals = append(goals, Goal { [2]GoalDef { eggsInRegion(Waterlands), birdsInRegion(Waterlands) }, rand.Intn(2) })

	goals = append(goals, Goal { [2]GoalDef { eggsInNest(Platform), nestsWithEggs(Platform) }, rand.Intn(2) })
	goals = append(goals, Goal { [2]GoalDef { eggsInNest(Cup), nestsWithEggs(Cup) }, rand.Intn(2) })
	goals = append(goals, Goal { [2]GoalDef { eggsInNest(Rocks), nestsWithEggs(Rocks) }, rand.Intn(2) })
	goals = append(goals, Goal { [2]GoalDef { eggsInNest(Canada), nestsWithEggs(Canada) }, rand.Intn(2) })

	goals = append(goals, Goal { [2]GoalDef { totalBirds(), sets() }, rand.Intn(2) })

	return goals
}

func (g *Game) chooseGoals() {
	g.goals = [4]Goal {}

	goals := allGoals()

	for round, idx := range rand.Perm(8)[4:] {
		g.goals[round] = goals[idx]
	}
}

func (g *Game) scoreGoals() {
	scores := make([]int, 0)
}