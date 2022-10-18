package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new label widget to show in the window.
	l, err := gtk.LabelNew("Hello, gotk3!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	// Add the label to the window.
	win.Add(l)

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func main() {
//	rand.Seed(time.Now().Unix())
//
//	game := new(Game)
//	game.init(5)
//
//	p := game.players[0]
//
//	pre := p.generatePregame()
//	pregame(p, game, pre[rand.Intn(len(pre))].a)
//
//	for _, b := range p.hand {
//		fmt.Println(b.name, b.cost)
//	}
//	fmt.Println(p.food)
//
//	fmt.Println(displayBirdArray(p.hand))
//
//	//for i := 0; i < 10; i++ {
//	//	m := p.chooseMove(game, 3)
//	//
//	//	moves := p.generateMoves(game)
//	//	a, b, c, d := 0, 0, 0, 0
//	//	for _, move := range moves {
//	//		switch move.t {
//	//		case PlayBird:
//	//			a++
//	//		case GainFood:
//	//			b++
//	//		case LayEggs:
//	//			c++
//	//		case DrawCards:
//	//			d++
//	//		}
//	//	}
//	//	fmt.Printf("Found %d %d %d %d\n", a, b, c, d)
//	//
//	//	switch m.t {
//	//	case PlayBird:
//	//		fmt.Println(m.t.String(), m.a.b, m.a.r, m.a.f)
//	//	case GainFood:
//	//		fmt.Println(m.t.String(), m.a.f)
//	//	case LayEggs:
//	//		fmt.Println(m.t.String(), m.a.e)
//	//	case DrawCards:
//	//		fmt.Println(m.t.String(), m.a.tray, m.a.ndeck)
//	//	}
//	//
//	//	m.f(p, game, m.a)
//	//
//	//	fmt.Println(p.String())
//	//}
//}
