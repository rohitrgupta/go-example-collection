package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")
	myWindow.Resize(fyne.NewSize(500, 500))

	text1 := canvas.NewText("Hello", color.Black)
	text2 := canvas.NewText("There", color.Black)
	text3 := canvas.NewText("(right)", color.Black)
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	text4 := canvas.NewText("centered", color.Black)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	myWindow.ShowAndRun()
}
