package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create the application
	myApp := app.New()
	myWindow := myApp.NewWindow("Fyne Button Ribbon")

	// Create buttons for the ribbon
	btnNew := widget.NewButton("New", func() {
		fmt.Println("New clicked")
	})
	btnOpen := widget.NewButton("Open", func() {
		fmt.Println("Open clicked")
	})
	btnSave := widget.NewButton("Save", func() {
		fmt.Println("Save clicked")
	})
	btnExit := widget.NewButton("Exit", func() {
		myApp.Quit()
	})

	// Arrange buttons horizontally
	ribbon := container.NewHBox(btnNew, btnOpen, btnSave, btnExit)

	// Example main content
	content := widget.NewLabel("Main application content goes here...")

	// Layout: ribbon at top, content below
	mainLayout := container.NewBorder(ribbon, nil, nil, nil, content)

	// Set content and show window
	myWindow.SetContent(mainLayout)
	myWindow.Resize(fyne.NewSize(400, 200))
	myWindow.ShowAndRun()
}
