package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	top := canvas.NewText("top bar", color.Black)
	left := canvas.NewText("left", color.Black)
	middle := canvas.NewText("content", color.Black)
	content := container.NewBorder(top, nil, left, nil, middle)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
