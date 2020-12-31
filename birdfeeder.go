package main

import "math/rand"

type Birdfeeder struct {
	diceIn []Dice
	diceOut []Dice
}

func (b *Birdfeeder) remove(d Dice) {
	var dice []Dice

	for _, i := range b.diceIn {
		if !(len(i.upface) == len(d.upface) && d.upface[0] == i.upface[0]) {
			dice = append(dice, i)
		}
	}

	b.diceIn = dice
	b.diceOut = append(b.diceOut, d)
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

func (d *Dice) hasFood(has Food) bool {
	for _, f := range d.upface {
		if f == has {
			return true
		}
	}
	return false
}