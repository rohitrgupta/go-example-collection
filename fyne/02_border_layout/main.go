package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	top := widget.NewLabel("top bar")
	left := widget.NewLabel("left")
	middle := widget.NewLabel("content")
	content := container.NewBorder(top, nil, left, nil, middle)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}
