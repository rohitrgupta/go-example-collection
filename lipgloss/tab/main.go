package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	width = 96
)

var (
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}
	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	tab = lipgloss.NewStyle().
		Border(tabBorder, true).
		BorderForeground(highlight).
		Padding(0, 1)

	activeTab = tab.Border(activeTabBorder, true)
	tabGap    = tab.
			BorderTop(false).
			BorderLeft(false).
			BorderRight(false)
)

func main() {
	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		activeTab.Render("Lip Gloss"),
		tab.Render("Blush"),
		tab.Render("Eye Shadow"),
		tab.Render("Mascara"),
		tab.Render("Foundation"),
	)

	gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap))
}
