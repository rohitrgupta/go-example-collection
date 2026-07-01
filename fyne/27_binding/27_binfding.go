package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create the application
	myApp := app.New()
	myWindow := myApp.NewWindow("Fyne Entry Binding Example")

	// Create a bound string
	boundStr := binding.NewString()
	err := boundStr.Set("Hello, Fyne!")
	if err != nil {
		log.Fatalf("Failed to set initial value: %v", err)
	}

	// Create an Entry bound to the string
	entry := widget.NewEntryWithData(boundStr)

	// Create a label bound to the same string
	label := widget.NewLabelWithData(boundStr)

	// Button to print the current bound value
	printBtn := widget.NewButton("Print Value", func() {
		val, err := boundStr.Get()
		if err != nil {
			log.Printf("Error reading bound value: %v", err)
			return
		}
		fmt.Println("Current value:", val)
	})

	// Layout
	content := container.NewVBox(
		widget.NewLabel("Type something:"),
		entry,
		widget.NewLabel("Live bound label:"),
		label,
		printBtn,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}
