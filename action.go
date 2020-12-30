package main

type ActionType int
const (
	Activated ActionType = iota
	Play
	Once
)

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
)

type Action struct {
	typ ActionType
	cause []Thing
	effect []Thing
}

type Thing struct {
	typ ThingType
	arg *Thing
}

func parse(s string) Action {

}