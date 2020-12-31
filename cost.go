package main

import (
	"fmt"
	"strings"
)

type Cost struct {
	cost []Component
}

type Component struct {
	options []Food
}

func readCost(s string) Cost {
	t := Cost { cost: make([]Component, 0) }

	if s == "" {
		return t
	}

	for _, component := range strings.Split(s, "+") {
		c := Component { options: make([]Food, 0) }

		for _, f := range strings.Split(component, "/") {
			c.options = append(c.options, Food(Atoi(f)))
		}

		t.cost = append(t.cost, c)
	}

	return t
}

func (c Cost) String() string {
	var br strings.Builder

	for a, i := range c.cost {
		for b, j := range i.options {
			fmt.Fprintf(&br, j.String())
			if b != len(i.options) - 1 {
				fmt.Fprintf(&br, "+")
			}
		}
		if a != len(c.cost) - 1 {
			fmt.Fprintf(&br, "/")
		}
	}
	fmt.Fprintf(&br, "\n")
	return br.String()
}