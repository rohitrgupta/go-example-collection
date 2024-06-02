package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	v1 := container.NewVBox(
		widget.NewLabel("Hello"),
		widget.NewButton("Exit", func() {
			w.Close()
		}),
	)
	v1.Add(widget.NewLabel("loram ipsum"))
	w.SetContent(v1)
	w.ShowAndRun()
}
