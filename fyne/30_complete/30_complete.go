package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	// Import fyne-x for CompletionEntry
	xwidget "fyne.io/x/fyne/widget"
)

func main() {
	// Create the application
	myApp := app.New()
	myWindow := myApp.NewWindow("CompletionEntry Example")

	// Sample data for auto-completion
	suggestions := []string{
		"Apple", "Apricot", "Avocado",
		"Banana", "Blackberry", "Blueberry",
		"Cherry", "Coconut", "Cranberry",
		"Grape", "Grapefruit", "Guava",
		"Mango", "Melon", "Mulberry",
		"Orange", "Papaya", "Peach",
		"Pear", "Pineapple", "Plum",
		"Strawberry", "Watermelon",
	}

	// Create a CompletionEntry
	entry := xwidget.NewCompletionEntry(suggestions)

	// Optional: Customize filtering behavior
	entry.OnChanged = func(s string) {
		// Filter suggestions dynamically
		var filtered []string
		for _, item := range suggestions {
			if len(s) == 0 || containsIgnoreCase(item, s) {
				filtered = append(filtered, item)
			}
		}
		entry.SetOptions(filtered)
		entry.ShowCompletion() // Show dropdown
	}

	// Add a label to show selected value
	selectedLabel := widget.NewLabel("Selected: (none)")

	// Handle when user selects an option
	entry.OnSubmitted = func(s string) {
		selectedLabel.SetText("Selected: " + s)
	}

	// Layout
	content := container.NewVBox(
		widget.NewLabel("Type a fruit name:"),
		entry,
		selectedLabel,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 200))
	myWindow.ShowAndRun()
}

// Helper: case-insensitive substring match
func containsIgnoreCase(str, substr string) bool {
	return len(str) >= len(substr) &&
		// Convert both to lowercase for comparison
		// (avoids importing strings.ToLower repeatedly)
		func() bool {
			s1 := []rune(str)
			s2 := []rune(substr)
			for i := range s1 {
				if i+len(s2) > len(s1) {
					break
				}
				match := true
				for j := range s2 {
					c1 := s1[i+j]
					c2 := s2[j]
					if c1 >= 'A' && c1 <= 'Z' {
						c1 += 'a' - 'A'
					}
					if c2 >= 'A' && c2 <= 'Z' {
						c2 += 'a' - 'A'
					}
					if c1 != c2 {
						match = false
						break
					}
				}
				if match {
					return true
				}
			}
			return false
		}()
}
