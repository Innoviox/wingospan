package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ActionType int
const (
	Activated ActionType = iota
	Play
	Once
)

func (a ActionType) String() string {
	return[...]string{"Activated","Play","Once"}[a]
}

type ThingType int
const (
	Cache ThingType = iota
	CacheOptional
	Hunt
	Tuck
	TuckFromDeck
	GainFromBirdfeeder
	GainAllFromBirdfeeder
	GainSupply
	PlaySecond
	Bonusify
	Draw
	DrawTray
	RepeatBrown
	RepeatHunter
	LayEgg
	LayEggAnother
	LayEggAny
	LayEggEach
	MoveIfRight
	RollOutside
	PlayersWithFewest
	AllPlayers
	Trade
	DiscardCard
	DiscardEgg
	DiscardEggAnother
	DiscardAtTurnEnd
	DiscardFood
	LayEggAction
	DiscardAnyEgg
	PlayBirdAction
	GainFoodAction
	SuccessfulHunt
)

func (t ThingType) String() string {
	return[...]string{"Cache","CacheOptional","Hunt","Tuck","TuckFromDeck","GainFromBirdfeeder","GainAllFromBirdfeeder","GainSupply","PlaySecond","Bonusify","Draw","DrawTray","RepeatBrown","RepeatHunter","LayEgg","LayEggAnother","LayEggAny","LayEggEach","MoveIfRight","RollOutside","PlayersWithFewest","AllPlayers","Trade","DiscardCard","DiscardEgg","DiscardEggAnother","DiscardAtTurnEnd","DiscardFood","LayEggAction","DiscardAnyEgg","PlayBirdAction","GainFoodAction","SuccessfulHunt"}[t]
}

type Action struct {
	typ ActionType
	cause *Things
	effect *Things
}

type Things struct {
	things []*Thing
	typ JoinType
}

type Thing struct {
	typ ThingType
	arg interface{}
}


func readAction(s string) *Action {
	act := new(Action)

	if s == "" {
		return act
	}

	parts := strings.Split(s, ",")

	act.typ = ActionType(strings.Index("APO", parts[0]))
	act.cause = parseThings(parts[1])
	if len(parts) == 3 {
		act.effect = parseThings(parts[2])
	}

	return act
}

func parseThings(s string) *Things {
	t := new(Things)

	if strings.Contains(s, "+") {
		t.typ = And
		for _, p := range strings.Split(s, "+") {
			t.things = append(t.things, parseThing(p))
		}
	} else if strings.Contains(s, ",") {
		t.typ = Or
		for _, p := range strings.Split(s, ",") {
			t.things = append(t.things, parseThing(p))
		}
	} else {
		t.typ = And
		t.things = append(t.things, parseThing(s))
	}

	return t
}

func parseThing(s string) *Thing {
	t := new(Thing)
	
	t.typ = ThingType(strings.Index("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdef", string(s[0])))

	if len(s) > 1 {
		arg := s[1:]
		i, _ := strconv.Atoi(arg)

		switch t.typ {
		case GainSupply, DiscardFood, GainFromBirdfeeder, Cache, RollOutside, GainFoodAction, GainAllFromBirdfeeder:
			t.arg = Food(i)
		case LayEggAny, LayEggAnother, LayEggEach:
			t.arg = Nest(i)
		case PlayersWithFewest, PlayBirdAction:
			t.arg = Region(i)
		case Hunt:
			t.arg = i
		case AllPlayers:
			t.arg = parseThing(arg)
		}
	}

	return t
}

func (a *Action) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "%s %s %s", a.typ.String(), a.cause.String(), a.effect.String())

	return br.String()
}

func (t *Things) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "%s ", t.typ.String())

	for _, i := range t.things {
		fmt.Fprintf(&br, "%s", i.String())
	}

	return br.String()
}

func (t *Thing) String() string {
	var br strings.Builder

	fmt.Fprintf(&br, "(%s ", t.typ.String())

	switch v := t.arg.(type) {
	case Food: fmt.Fprintf(&br, v.String()) // why can't i do the thing :(
	case Nest: fmt.Fprintf(&br, v.String())
	case Region: fmt.Fprintf(&br, v.String())
	case int: fmt.Fprintf(&br, "%d", v)
	case *Thing: fmt.Fprintf(&br, v.String())
	}

	fmt.Fprintf(&br, ")")

	return br.String()
}