package main

// A simple program that counts down from 5 and then exits.

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	PaddingTop(1).
	PaddingBottom(1).
	PaddingLeft(4).
	PaddingRight(4).
	Align(lipgloss.Center).
	BorderStyle(lipgloss.NormalBorder()).
	Width(22)

func main() {
	p := tea.NewProgram(model(5))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model int

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return style.Render("Hello World")
}
