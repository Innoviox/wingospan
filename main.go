package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/innoviox/wingospan/gui"
	"github.com/innoviox/wingospan/model"
)

func main() {
	game := new(model.Game)
	game.Init(1)

	p := game.Players[0]

	a := app.New()
	w := a.NewWindow("Wingspan")

	w.SetContent(gui.DisplayBirds(p.Hand))

	w.ShowAndRun()

	//game.init(1)
	//
	//p := game.players[0]
	//
	//pre := p.generatePregame()
	//pregame(p, game, pre[rand.Intn(len(pre))].a)
	//
	//for _, b := range p.hand {
	//	fmt.Println(b.name, b.cost)
	//}
	//fmt.Println(p.food)
	//
	//fmt.Println(displayBirdArray(p.hand))
	//
	//game.start()
}
