package main

import "math/rand"

type Birdfeeder struct {
	diceIn []Dice
	diceOut []Dice
}

type Dice struct {
	upface []Food
}

var (
	faces = [6][]Food { { Worm, Seed }, { Worm }, { Seed }, { Fish }, { Rodent }, { Berry }}
)

func Die() Dice {
	d := Dice {}
	d.roll()
	return d
}

func (d *Dice) roll() {
	d.upface = faces[rand.Intn(len(faces))]
}