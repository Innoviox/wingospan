package main

type Board struct {
	rows [3][]Bird
	r_idxs [3]int
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

func (b *Board) playBird(bird Bird, r Region) {
	b.rows[r] = append(b.rows[r], bird)
	b.r_idxs[r]++
}