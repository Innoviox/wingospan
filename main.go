/*
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	game := new(Game)
	game.init(1)

	p := game.players[0]

	pre := p.generatePregame()
	pregame(p, game, pre[rand.Intn(len(pre))].a)

	for _, b := range p.hand {
		fmt.Println(b.name, b.cost)
	}
	fmt.Println(p.food)

	fmt.Println(displayBirdArray(p.hand))

	game.start()
}
*/
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
