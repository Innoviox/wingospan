package main

type Bird struct { // todo store name types for bonus; images
	name string
	cost Cost

	points int
	nest Nest
	eggLimit int
	eggs int
	action Action
}