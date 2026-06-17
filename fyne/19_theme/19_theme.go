package main

import (
	"image/color"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// MyTheme implements fyne.Theme interface
type MyTheme struct{}

// Color overrides default theme colors
func (m MyTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

// Font overrides default fonts (return nil to use default)
func (m MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

// Icon overrides default icons
func (m MyTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// Size overrides default sizes
func (m MyTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 18 // Larger text
	case theme.SizeNamePadding:
		return 12
	default:
		return theme.DefaultTheme().Size(name)
	}
}

func main() {
	a := app.New()
	a.Settings().SetTheme(&MyTheme{}) // Apply custom theme

	w := a.NewWindow("Fyne Custom Theme Example")
	w.Resize(fyne.NewSize(400, 300))

	label := widget.NewLabel("Hello, Custom Theme!")
	button := widget.NewButton("Click Me", func() {
		log.Println("Button clicked!")
	})

	rect := canvas.NewRectangle(color.NRGBA{R: 255, G: 87, B: 34, A: 255})
	rect.SetMinSize(fyne.NewSize(100, 50))

	content := container.NewVBox(label, button, rect)
	w.SetContent(content)
	go func() {
		for {
			time.Sleep(50 * time.Millisecond)
			if label.Visible() {
				break
			}
		}
		fyne.CurrentApp().Settings().SetTheme(myTheme.ToFyneTheme())
	}()

	w.ShowAndRun()
}
