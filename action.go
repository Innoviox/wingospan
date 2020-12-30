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
	LayEggAction
	DiscardAnyEgg
	PlayBirdAction
	GainFoodAction
	SuccessfulHunt
)

type JoinType int
const (
	And JoinType = iota
	Or
)

type Action struct {
	typ ActionType
	cause Things
	effect Things
}

type Things struct {
	things []Thing
	typ JoinType
}

type Thing struct {
	typ ThingType
	arg *interface{}
}

func parse(s string) Action {

}