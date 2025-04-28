package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(1).
		PaddingBottom(1).
		PaddingLeft(4).
		PaddingRight(4).
		Align(lipgloss.Right).
		BorderStyle(lipgloss.NormalBorder()).
		Width(22)

	var style2 = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(0).
		PaddingBottom(0).
		PaddingLeft(3).
		PaddingRight(3).
		Align(lipgloss.Right).
		BorderStyle(lipgloss.NormalBorder()).
		Width(22)

	// fmt.Println(style.Render("Hello,\nkitty")
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Bottom, style.Render("Hello,\nkitty"), style2.Render("Hello,\nkitty")))
}
