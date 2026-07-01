package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new Fyne application
	myApp := app.New()
	myWindow := myApp.NewWindow("Fyne Menu Example")

	// --- Main Menu ---
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("New", func() {
			dialog.ShowInformation("Menu Action", "New File clicked", myWindow)
		}),
		fyne.NewMenuItem("Copy", func() {
			dialog.ShowInformation("Menu Action", "Copy clicked", myWindow)
		}),
		fyne.NewMenuItem("Edit", func() {
			dialog.ShowInformation("Menu Action", "Edit clicked", myWindow)
		}),
		fyne.NewMenuItem("Save", func() {
			dialog.ShowInformation("Menu Action", "Save clicked", myWindow)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Quit", func() {
			myApp.Quit()
		}),
	)

	editMenu := fyne.NewMenu("Maps",
		fyne.NewMenuItem("Bitmap", func() {
			dialog.ShowInformation("Menu Action", "Bitmap clicked", myWindow)
		}),
		fyne.NewMenuItem("Bitmap Action", func() {
			dialog.ShowInformation("Menu Action", "Bitmap Action clicked", myWindow)
		}),
		fyne.NewMenuItem("Modbus", func() {
			dialog.ShowInformation("Menu Action", "Modbus clicked", myWindow)
		}),
		fyne.NewMenuItem("Modbus Action", func() {
			dialog.ShowInformation("Menu Action", "Modbus Action clicked", myWindow)
		}),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowInformation("About", "Fyne Menu Example v1.0", myWindow)
		}),
	)

	// Set the main menu
	myWindow.SetMainMenu(fyne.NewMainMenu(fileMenu, editMenu, helpMenu))

	// --- Context Menu (Right Click) ---
	// contextMenu := fyne.NewMenu("",
	// 	fyne.NewMenuItem("Option 1", func() {
	// 		dialog.ShowInformation("Context Menu", "Option 1 clicked", myWindow)
	// 	}),
	// 	fyne.NewMenuItem("Option 2", func() {
	// 		dialog.ShowInformation("Context Menu", "Option 2 clicked", myWindow)
	// 	}),
	// )

	label := widget.NewLabel("Right-click here for context menu")
	// label.OnTappedSecondary = func(ev *fyne.PointEvent) {
	// 	widget.ShowPopUpMenuAtPosition(contextMenu, myWindow.Canvas(), ev.AbsolutePosition)
	// }

	// Set content
	myWindow.SetContent(container.NewVBox(
		label,
		widget.NewLabel("Use the menu bar above or right-click the label."),
	))

	myWindow.Resize(fyne.NewSize(400, 200))
	myWindow.ShowAndRun()
}
