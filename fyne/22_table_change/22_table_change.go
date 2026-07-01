package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne Table Remove Row")

	// Sample data
	data := [][]string{
		{"Alice", "25"},
		{"Bob", "30"},
		{"Charlie", "35"},
	}

	// Table creation
	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			cell.(*widget.Label).SetText(data[id.Row][id.Col])
		},
	)

	// Button to remove the second row (index 1)
	removeBtn := widget.NewButton("Remove Row 2", func() {
		if len(data) > 1 {
			// Remove row at index 1
			data = append(data[:1], data[2:]...)
			table.Refresh() // Refresh table to reflect changes
		}
	})

	w.SetContent(container.NewBorder(nil, removeBtn, nil, nil, table))
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
