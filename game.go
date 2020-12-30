package main

type Game struct {
	players []Player

	deck []Bird
	market []Bird

	birdfeeder Birdfeeder

	goals [4]Goal
}
