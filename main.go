package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	game := new(Game)

	// readAction birds
	game.deck = make([]Bird, 0)

	file, _ := os.Open("parse/birds.csv")
	r := csv.NewReader(file)

	records, _ := r.ReadAll()

	for i, line := range records {
		if i == 0 {
			continue
		}

		game.deck = append(game.deck, Bird {
			name:     line[0],
			region:   parseRegion(line[1]),
			cost:     readCost(line[2]),
			points:   Atoi(line[3]),
			nest:     Nest(Atoi(line[4])),
			eggLimit: Atoi(line[5]),
			wingspan: Atoi(line[6]),
			action:   readAction(line[8]),
		})
	}

	for _, bird := range game.deck {
		fmt.Println(bird.action)
	}
}