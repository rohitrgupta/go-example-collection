package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	text1 := widget.NewLabel("Hello")
	text2 := widget.NewLabel("There")
	text2.Move(fyne.NewPos(20, 20))
	content := container.NewWithoutLayout(text1, text2)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}
