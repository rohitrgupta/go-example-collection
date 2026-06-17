package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

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
	a := app.New()

	// Apply custom theme
	a.Settings().SetTheme(&SmallPaddingTheme{Theme: theme.DefaultTheme()})

	w := a.NewWindow("Reduced Widget Padding")

	content := container.NewVBox(
		widget.NewLabel("Label 1"),
		widget.NewButton("Button 1", nil),
		widget.NewLabel("Label 2"),
		widget.NewButton("Button 2", nil),
		widget.NewEntry(),
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
