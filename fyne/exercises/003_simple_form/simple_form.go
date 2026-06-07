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
	b := widget.NewButton("Click", func() {
		app.output.SetText(e.Text)
	})
	vb := container.NewVBox(app.output, e, b)
	return vb
}
