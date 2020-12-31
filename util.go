package main

import "strconv"

type Food int
const (
	Worm Food = iota
	Seed
	Fish
	Rodent
	Berry
	Any
)

type Nest int // todo
const (
	Platform Nest = iota
	Cup
	Rocks
	Canada
	Star
)

type Region int
const ( // todo right names
	Forest Region = iota
	Grasslands
	Waterlands
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}