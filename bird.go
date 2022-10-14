package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Bird struct { // todo store name types for bonus; images
	name   string
	region []Region
	cost   Cost

	points   int
	nest     Nest
	eggLimit int
	eggs     int
	wingspan int

	caches int
	tucks  int

	action *Action
}

func (b *Bird) activateBrown(g *Game, p *Player) {
	if b.action.typ != Activated {
		return
	}

	if b.action == nil || b.action.cause == nil {
		return
	}

	if b.action.effect == nil {
		// then the action is stored in the cause
	} else {
		switch b.action.cause.typ {
		case And:
			for _, t := range b.action.cause.things {
				t.activateThing(g, p, b)
			}
			break
		case Or:
			//choice := p.promptChoice(b.action.cause.things)
		}
	}
}

func (t *Thing) activateThing(g *Game, p *Player, b *Bird) {
	//switch t.typ {
	//case Draw:
	//	//card = p.promptDraw()
	//
	//case LayEgg:
	//	...
	//case LayEggAnother:
	//	...
	//case LayEggEach:
	//	...
	//}
}

func (b *Bird) StringFor(i int) string {
	var br strings.Builder
	switch i {
	case 0:
		return b.name
	case 1:
		fmt.Fprintf(&br, "[")
		for j := 0; j < len(b.region); j++ {
			fmt.Fprintf(&br, b.region[j].String(), ", ")
		}
		fmt.Fprintf(&br, "]")

		return br.String()
	case 2:
		return b.cost.String()
	case 3:
		return strconv.Itoa(b.points)
	case 4:
		return b.nest.String()
	case 5:
		return fmt.Sprintf("%d/%d", b.eggs, b.eggLimit)
	case 6:
		return strconv.Itoa(b.wingspan)
	case 7:
		return fmt.Sprintf("%d, %d", b.caches, b.tucks)
	case 8:
		return b.action.String()
	}
	return ""
}
