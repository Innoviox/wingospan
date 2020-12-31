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

type JoinType int
const (
	And JoinType = iota
	Or
)


func Atoi(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

func splitString[type T] (s string, parse func(string) *T) (JoinType, []*T) {
	typ := And
	things := make([]*T, 0)

	if strings.Contains(s, ";") {
		typ = And
		for _, p := range strings.Split(s, ";") {
			things = append(things, parse(p))
		}
	} else if strings.Contains(s, ",") {
		typ = Or
		for _, p := range strings.Split(s, ",") {
			things = append(things, parse(p))
		}
	} else {
		typ = And
		things = append(things, parse(s))
	}

	return (typ, things)
}
