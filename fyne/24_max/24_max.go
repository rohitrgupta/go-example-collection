package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// PageManager handles switching between pages
type PageManager struct {
	mainContainer *fyne.Container
}

func (pm *PageManager) Show(page fyne.CanvasObject) {
	pm.mainContainer.Objects = []fyne.CanvasObject{page}
	pm.mainContainer.Refresh()
}

func main() {
	a := app.New()
	w := a.NewWindow("Fyne Multi-Page Example")
	w.Resize(fyne.NewSize(400, 300))

	// Main container that will hold the active page
	mainContainer := container.NewStack()
	pm := &PageManager{mainContainer: mainContainer}

	// Declare pages first (nil for now)
	var page1, page2, page3 fyne.CanvasObject

	// Create pages
	page1 = container.NewVBox(
		widget.NewLabel("This is Page 1"),
		widget.NewButton("Go to Page 2", func() {
			pm.Show(page2)
		}),
	)

	page2 = container.NewVBox(
		widget.NewLabel("This is Page 2"),
		widget.NewButton("Go to Page 3", func() {
			pm.Show(page3)
		}),
		widget.NewButton("Back to Page 1", func() {
			pm.Show(page1)
		}),
	)

	page3 = container.NewVBox(
		widget.NewLabel("This is Page 3"),
		widget.NewButton("Back to Page 2", func() {
			pm.Show(page2)
		}),
	)

	// Start with Page 1
	pm.Show(page1)

	// Set window content
	w.SetContent(mainContainer)
	w.ShowAndRun()
}
