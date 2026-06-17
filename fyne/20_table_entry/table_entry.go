package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Table dimensions
const rows = 5
const cols = 3

// SmallPaddingTheme overrides default theme sizes
type SmallPaddingTheme struct {
	fyne.Theme
}

func (s SmallPaddingTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 2 // default is 4
	case theme.SizeNameInnerPadding:
		return 1 // default is 2
	default:
		return s.Theme.Size(name)
	}
}

func (s SmallPaddingTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return s.Theme.Color(name, variant)
}

func (s SmallPaddingTheme) Font(style fyne.TextStyle) fyne.Resource {
	return s.Theme.Font(style)
}
func main() {
	// Create the application
	myApp := app.New()
	myApp.Settings().SetTheme(&SmallPaddingTheme{Theme: theme.DefaultTheme()})

	myWindow := myApp.NewWindow("Fyne Table with Entry")

	// Data storage for table entries
	data := make([][]string, rows)
	for i := range data {
		data[i] = make([]string, cols)
		for j := range data[i] {
			data[i][j] = fmt.Sprintf("R%dC%d", i+1, j+1) // initial values
		}
	}

	// Create the table
	table := widget.NewTable(
		// Size function
		func() (int, int) {
			return rows, cols
		},
		// Create template cell
		func() fyne.CanvasObject {
			entry := widget.NewEntry()
			entry.SetPlaceHolder("Enter text")
			return entry
		},
		// Update cell content
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			entry := obj.(*widget.Entry)
			entry.SetText(data[id.Row][id.Col])

			// Update data when entry changes
			entry.OnChanged = func(s string) {
				data[id.Row][id.Col] = s
			}
		},
	)

	// Optional: Set column widths
	for c := 0; c < cols; c++ {
		table.SetColumnWidth(c, 120)
	}

	// Layout
	content := container.NewBorder(nil, nil, nil, nil, table)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
