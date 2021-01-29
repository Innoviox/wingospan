package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Food int
const (
	Worm Food = iota
	Seed
	Fish
	Rodent
	Berry
	Any
)

func (f Food) String() string {
	return [...]string{"Worm", "Seed", "Fish", "Rodent", "Berry", "Any"}[f]
}

type Nest int // todo
const (
	Platform Nest = iota
	Cup
	Rocks
	Canada
	Star
	None
)

func (n Nest) String() string {
	return[...]string{"Platform","Cup","Rocks","Canada","Star","None"}[n]
}

type Region int
const ( // todo right names
	Forest Region = iota
	Grasslands
	Waterlands
)

func (r Region) String() string {
	return[...]string{"Forest","Grasslands","Waterlands"}[r]
}

type JoinType int
const (
	And JoinType = iota
	Or
)

func (j JoinType) String() string {
	return[...]string{"And","Or",")"}[j]
}

type MoveType int
const (
	PlayBird MoveType = iota
	GainFood
	LayEggs
	DrawCards
	PreGame
)

func (m MoveType) String() string {
	return[...]string{"PlayBird","GainFood","LayEggs","DrawCards","PreGame"}[m]
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

func combrep(n int, lst []string) [][]string {
	if n == 0 {
		return [][]string{nil}
	}
	if len(lst) == 0 {
		return nil
	}
	r := combrep(n, lst[1:])
	for _, x := range combrep(n-1, lst) {
		r = append(r, append(x, lst[0]))
	}
	return r
}

func parseRegion(s string) []Region {
	r := make([]Region, 0)
	for _, i := range strings.Split(s, "/") {
		r = append(r, Region(Atoi(i)))
	}
	return r
}

type Eggs [][3]int // a list of (region, birdn, #eggs)

// I FUCKING HATE YOU GO. WHAT FUCKING LANGUAGE DOESNT FUCKING HAVE GENERICS. FICK YOU
func cloneBirds(arr []Bird) []Bird {
	deck := make([]Bird, 0)
	for i := 0; i < len(arr); i++ {
		deck = append(deck, arr[i].clone())
	}
	return deck
}

func cloneDice(arr []Dice) []Dice {
	deck := make([]Dice, 0)
	for i := 0; i < len(arr); i++ {
		deck = append(deck, arr[i].clone())
	}
	return deck
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func mapToString(m map[Food]int) string {
	b := new(bytes.Buffer)
	for i := 0; i < 5; i++ {
		f := Food(i)
		fmt.Fprintf(b, "%s=\"%d\",", f.String(), m[f])
	}
	fmt.Fprintf(b, "\n")
	return b.String()
}