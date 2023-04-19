package model

import (
	"github.com/innoviox/wingospan/model/util"
	"math/rand"
	"sort"
)

type GoalDef func(*Player) int

var (
	points = [4][5]int{
		{4, 1, 0, 0, 0},
		{5, 2, 1, 0, 0},
		{6, 3, 2, 0, 0},
		{7, 4, 3, 0, 0},
	}
)

type Goal struct {
	sides  [2]GoalDef
	chosen int
}

func (g Goal) upface() GoalDef {
	return g.sides[g.chosen]
}

func eggsInRegion(region util.Region) GoalDef {
	return func(p *Player) int {
		total := 0

		for _, bird := range p.board.rows[region] {
			total += bird.eggs
		}

		return total
	}
}

func birdsInRegion(region util.Region) GoalDef {
	return func(p *Player) int {
		total := 0

		for range p.board.rows[region] {
			total++
		}

		return total
	}
}

func eggsInNest(nest util.Nest) GoalDef {
	return func(p *Player) int {
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

func nestsWithEggs(nest util.Nest) GoalDef {
	return func(p *Player) int {
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
	return func(p *Player) int {
		total := 0

		for _, row := range p.board.rows {
			for range row {
				total++
			}
		}

		return total
	}
}

func sets() GoalDef {
	return func(p *Player) int {
		return 0 // todo
	}
}

func allGoals() []Goal {
	goals := make([]Goal, 0)

	goals = append(goals, Goal{[2]GoalDef{eggsInRegion(util.Forest), birdsInRegion(util.Forest)}, rand.Intn(2)})
	goals = append(goals, Goal{[2]GoalDef{eggsInRegion(util.Grasslands), birdsInRegion(util.Grasslands)}, rand.Intn(2)})
	goals = append(goals, Goal{[2]GoalDef{eggsInRegion(util.Waterlands), birdsInRegion(util.Waterlands)}, rand.Intn(2)})

	goals = append(goals, Goal{[2]GoalDef{eggsInNest(util.Platform), nestsWithEggs(util.Platform)}, rand.Intn(2)})
	goals = append(goals, Goal{[2]GoalDef{eggsInNest(util.Cup), nestsWithEggs(util.Cup)}, rand.Intn(2)})
	goals = append(goals, Goal{[2]GoalDef{eggsInNest(util.Rocks), nestsWithEggs(util.Rocks)}, rand.Intn(2)})
	goals = append(goals, Goal{[2]GoalDef{eggsInNest(util.Canada), nestsWithEggs(util.Canada)}, rand.Intn(2)})

	goals = append(goals, Goal{[2]GoalDef{totalBirds(), sets()}, rand.Intn(2)})

	return goals
}

func (g *Game) chooseGoals() {
	g.goals = [4]Goal{}

	goals := allGoals()

	for round, idx := range rand.Perm(8)[4:] {
		g.goals[round] = goals[idx]
	}
}

func (g *Game) scoreGoals() {
	var scores = map[int][]*Player{}

	goal := g.goals[g.round].upface()
	for _, p := range g.Players {
		val := goal(p)
		arr, in := scores[val]

		if in {
			scores[val] = append(arr, p)
		} else {
			scores[val] = []*Player{p}
		}
	}

	var order sort.IntSlice = make([]int, len(scores))
	i := 0
	for k := range scores {
		order[i] = k
		i++
	}
	sort.Sort(order)

	j := 0
	for _, o := range order {
		for _, p := range scores[o] {
			p.score += points[g.round][j]
			j++
		}
	}
}
