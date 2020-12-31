package main

import (
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

func Atoi(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

func parseRegion(s string) []Region {
	r := make([]Region, 0)
	for _, i := range strings.Split(s, "/") {
		r = append(r, Region(Atoi(i)))
	}
	return r
}

type Eggs [][]int // a list of (region, birdn, #eggs)