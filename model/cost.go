package model

import (
	"fmt"
	"github.com/innoviox/wingospan/model/util"
	"strings"
)

type Cost struct {
	cost []Component
}

type Component struct {
	options []util.Food
}

func readCost(s string) Cost {
	t := Cost{cost: make([]Component, 0)}

	if s == "" {
		return t
	}

	for _, component := range strings.Split(s, "+") {
		c := Component{options: make([]util.Food, 0)}

		for _, f := range strings.Split(component, "/") {
			c.options = append(c.options, util.Food(util.Atoi(f)))
			if f == "5" {
				c.options = []util.Food{util.Worm, util.Seed, util.Fish, util.Rodent, util.Berry}
			}
		}

		t.cost = append(t.cost, c)
	}

	return t
}

func (c Cost) options() [][]util.Food {
	opts := make([][]util.Food, 0)

	for i, component := range c.cost {
		new := make([][]util.Food, 0)
		for _, o := range component.options {
			if i == 0 {
				opts = append(opts, []util.Food{o})
			} else {
				for _, opt := range opts {
					new = append(new, append(opt, o))
				}
			}
		}

		if len(new) != 0 {
			opts = new
		}
	}

	return opts
}

func (c Cost) String() string {
	var br strings.Builder

	for a, i := range c.cost {
		for b, j := range i.options {
			fmt.Fprintf(&br, j.String())
			if b != len(i.options)-1 {
				fmt.Fprintf(&br, "/")
			}
		}
		if a != len(c.cost)-1 {
			fmt.Fprintf(&br, "+")
		}
	}
	//fmt.Fprintf(&br, "\n")
	return br.String()
}
