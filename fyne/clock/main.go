package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	t := time.Now().Format("15:04:05")
	clock.SetText(t)
}

func tickerThread(clock *widget.Label) {
	for range time.Tick(time.Second) {
		updateTime(clock)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Clock")
	clock := widget.NewLabel("")
	updateTime(clock)
	w.SetContent(clock)
	go tickerThread(clock)
	w.ShowAndRun()
}
