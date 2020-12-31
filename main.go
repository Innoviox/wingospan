package main

import (
	"encoding/csv"
	"os"
)

func main() {
	game := new(Game)

	// readAction birds
	game.deck = make([]Bird, 0)

	file, _ := os.Open("readAction/birds.csv")
	r := csv.NewReader(file)

	records, _ := r.ReadAll()
	for _, line := range records {
		game.deck = append(game.deck, Bird {
			name:     line[0],
			region:   Region(Atoi(line[1])),
			cost:     readCost(line[2]),
			points:   Atoi(line[3]),
			nest:     Nest(Atoi(line[4])),
			eggLimit: Atoi(line[5]),
			eggs:     0,
			wingspan: Atoi(line[6]),
			action:   readAction(line[8]),
		})
	}
}