package model

import (
	"fmt"
	"github.com/innoviox/wingospan/model/util"
	"strings"
)

type Bird struct { // todo store Name types for bonus; images
	Name   string
	region []util.Region
	Cost   Cost

	points   int
	nest     util.Nest
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
		case util.And:
			for _, t := range b.action.cause.things {
				t.activateThing(g, p, b)
			}
			break
		case util.Or:
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
	switch i {
	case 0:
		return b.Name
	case 1:
		return strings.Join(util.ArrToString(b.region), ", ")
	case 2:
		return b.Cost.String()
	case 3:
		return fmt.Sprintf("Points: %d", b.points)
	case 4:
		return b.nest.String()
	case 5:
		return fmt.Sprintf("Eggs: %d/%d", b.eggs, b.eggLimit)
	case 6:
		return fmt.Sprintf("Wingspan: %d", b.wingspan)
	case 7:
		return fmt.Sprintf("Caches: %d, Tucks: %d", b.caches, b.tucks)
	case 8:
		return b.action.String()
	}
	return ""
}
