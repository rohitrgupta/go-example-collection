package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.SetContent(widget.NewLabel("Hello"))
	w.SetMaster()
	w.Show()
	w2 := a.NewWindow("Hello World 2")
	w2.SetContent(widget.NewLabel("Hello 2"))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()
	a.Run()
}
