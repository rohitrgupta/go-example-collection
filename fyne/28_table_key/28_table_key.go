package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// CustomEntry wraps widget.Entry to capture special key events
type CustomEntry struct {
	widget.Entry
	row, col int
}

func NewCustomEntry(row, col int) *CustomEntry {
	e := &CustomEntry{row: row, col: col}
	e.ExtendBaseWidget(e)
	return e
}

// KeyDown handles special key presses
func (e *CustomEntry) KeyDown(ev *fyne.KeyEvent) {
	switch ev.Name {
	case fyne.KeyReturn, fyne.KeyEnter:
		fmt.Printf("Enter pressed at cell (%d, %d)\n", e.row, e.col)
	case fyne.KeyTab:
		fmt.Printf("Tab pressed at cell (%d, %d)\n", e.row, e.col)
	case fyne.KeyLeft:
		fmt.Printf("Left arrow at cell (%d, %d)\n", e.row, e.col)
	case fyne.KeyRight:
		fmt.Printf("Right arrow at cell (%d, %d)\n", e.row, e.col)
	case fyne.KeyUp:
		fmt.Printf("Up arrow at cell (%d, %d)\n", e.row, e.col)
	case fyne.KeyDown:
		fmt.Printf("Down arrow at cell (%d, %d)\n", e.row, e.col)
	default:
		// Pass other keys to default handler
		e.Entry.KeyDown(ev)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Fyne Table Special Key Handling")

	// Create a 5x3 table
	table := widget.NewTable(
		func() (int, int) { return 5, 3 },
		func() fyne.CanvasObject {
			// Each cell will be a CustomEntry
			return NewCustomEntry(0, 0)
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			// Update cell coordinates
			if ce, ok := obj.(*CustomEntry); ok {
				ce.row = id.Row
				ce.col = id.Col
				ce.SetText(fmt.Sprintf("R%dC%d", id.Row, id.Col))
			}
		},
	)

	// Optional: Add a background to see table boundaries
	bg := canvas.NewRectangle(color.NRGBA{R: 220, G: 220, B: 220, A: 255})
	content := container.NewMax(bg, table)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}
