package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// show different labels each with differen style

func main() {
	a := app.New()
	win := a.NewWindow("Text Style")
	s := fyne.TextStyle{Bold: true, Italic: true}
	l := widget.NewLabelWithStyle("Hello World!!", fyne.TextAlignCenter, s)
	l.Importance = widget.HighImportance
	win.SetContent(l)
	win.ShowAndRun()
}
