package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var data = [][]string{[]string{"top left", "top right"},
	[]string{"bottom left", "bottom right"}}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Table Widget")

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})
	title := widget.NewLabel("Site")
	intro := widget.NewLabel("Start")

	top := container.NewHBox(title, widget.NewSeparator(), intro)

	nb := container.NewBorder(top, nil, nil, nil, list)

	myWindow.SetContent(nb)
	myWindow.Resize(fyne.NewSize(1200, 600))

	myWindow.ShowAndRun()
}
