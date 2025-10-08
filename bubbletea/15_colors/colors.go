package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	// Define the number of colors per row for output formatting
	const colorsPerRow = 16

	// Loop through all 256 indexed colors
	for i := 0; i < 256; i++ {
		// Create a style with the current color index as foreground
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(fmt.Sprintf("%d", i)))

		// Print the color index styled
		fmt.Printf("%s ", style.Render(fmt.Sprintf("%3d", i)))

		style = lipgloss.NewStyle().Background(lipgloss.Color(fmt.Sprintf("%d", i)))

		// Print the color index styled
		fmt.Printf("%s ", style.Render(fmt.Sprintf("%3d", i)))

		// New line after certain number of colors per row
		if (i+1)%colorsPerRow == 0 {
			fmt.Println()
		}
	}
}
