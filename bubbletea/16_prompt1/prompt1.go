// Package main demonstrates how promptkit/selection is used.
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/erikgeiser/promptkit/selection"
)

type model struct {
	sp     *selection.Selection[string]
	choice string
}

func New() model {
	sp := selection.New("What do you pick?", []string{"Horse", "Car", "Plane", "Bike"})
	sp.PageSize = 3

	return model{
		sp: sp,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		return fmt.Sprintf("You picked: %s\n", m.choice)
	}
	return ""
}

func main() {
	sp := selection.New("What do you pick?", []string{"Horse", "Car", "Plane", "Bike"})
	sp.PageSize = 3

	choice, err := sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the final choice
	_ = choice
}
