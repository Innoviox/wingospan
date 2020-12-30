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
)