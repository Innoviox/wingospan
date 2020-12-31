package main

type Cost struct {
	cost []Component
}

type Component struct {
	options []Food
	typ JoinType
}

func readCost(s string) Cost {

}