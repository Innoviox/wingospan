package gui

import (
	"fyne.io/fyne/v2"
	container2 "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/innoviox/wingospan/model"
)

func DisplayBird(b model.Bird) *fyne.Container {
	container := container2.NewVBox()

	container.Add(widget.NewLabel(b.Name))
	container.Add(widget.NewLabel(b.Cost.String()))

	return container
}

func DisplayBirds(birds []model.Bird) *fyne.Container {
	container := container2.NewHBox()

	for _, b := range birds {
		container.Add(DisplayBird(b))
	}

	return container
}
