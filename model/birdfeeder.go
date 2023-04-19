package model

import (
	"github.com/innoviox/wingospan/model/util"
	"math/rand"
)

type Birdfeeder struct {
	diceIn  []Dice
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
	upface []util.Food
}

var (
	faces = [6][]util.Food{{util.Worm, util.Seed}, {util.Worm}, {util.Seed}, {util.Fish}, {util.Rodent}, {util.Berry}}
)

func Die() Dice {
	d := Dice{}
	d.roll()
	return d
}

func (d *Dice) roll() {
	d.upface = faces[rand.Intn(len(faces))]
}

func (d *Dice) hasFood(has util.Food) bool {
	for _, f := range d.upface {
		if f == has {
			return true
		}
	}
	return false
}
