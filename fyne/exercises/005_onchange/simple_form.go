package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
}

var MyApp App

func main() {
	a := app.New()
	win := a.NewWindow("Form")
	content := MyApp.MakeUi()
	win.SetContent(content)
	win.ShowAndRun()
}

func (app *App) MakeUi() *fyne.Container {
	app.output = widget.NewLabel("Hello")
	e := widget.NewEntry()
	e.OnChanged = func(value string) {
		app.output.SetText(value)
	}
	vb := container.NewVBox(app.output, e)
	return vb
}
