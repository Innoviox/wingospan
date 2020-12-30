package main

type Food int
const (
	Worm Food = iota
	Seed
	Fish
	Rodent
	Berry
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

