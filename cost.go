package main

import "strings"

type Cost struct {
	cost []Component
}

type Component struct {
	options []Food
}

func readCost(s string) Cost {
	t := Cost { cost: make([]Component, 0) }

	for _, component := range strings.Split(s, "+") {
		c := Component { options: make([]Food, 0) }

		for _, f := range strings.Split(component, "/") {
			c.options = append(c.options, Food(Atoi(f)))
		}

		t.cost = append(t.cost, c)
	}

	return t
}