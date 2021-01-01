package main

type Bird struct { // todo store name types for bonus; images
	name string
	region []Region
	cost Cost

	points int
	nest Nest
	eggLimit int
	eggs int
	wingspan int

	caches int
	tucks int

	action *Action
}

func (b *Bird) activateBrown(g *Game, p *Player) {
	if b.action.typ != Activated {
		return
	}

	switch b.action.cause.typ {
	case And:
		for _, t := range b.action.cause.things {

		}
	}
}