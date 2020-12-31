package main

type Board struct {
	rows [3][]Bird
}

func (b *Board) rawScore() int {
	s := 0

	for _, row := range b.rows {
		for _, bird := range row {
			s += bird.points + bird.eggs + bird.caches + bird.tucks
		}
	}

	return s
}

