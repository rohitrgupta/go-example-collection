package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	win := a.NewWindow("Hello")
	l := widget.NewLabel("Hello World!!")
	win.SetContent(l)
	win.ShowAndRun()
}
