package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func isIntValidator(s string) error {

	// chack if the string is an integer
	if _, err := strconv.Atoi(s); err != nil {
		return fmt.Errorf("Value is not an integer")
	}
	return nil
}

const (
	firstNameField = iota
	lastNameField
	emailField
	phoneField
	stateField
	businessField
	totalFields
)

var (
	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F5F83B")).
			Background(lipgloss.Color("#000000")).
			Width(12)

	inputStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#4A90D9")).
			Foreground(lipgloss.Color("#EAF4FF"))

	focusedInputStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#4A90D9")).
				Foreground(lipgloss.Color("#111111"))

	stateStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#4A90D9")).
			Foreground(lipgloss.Color("#CFE8FF"))

	focusedStateStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#4A90D9")).
				Background(lipgloss.Color("#000000"))

	checkboxStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#4A90D9")).
			Width(1)

	focusedCheckboxStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#7CB6ED")).
				Foreground(lipgloss.Color("#111111"))

	containerStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#000000")).
			Padding(1, 2)
)

type model struct {
	inputs    []textinput.Model
	focus     int
	states    []string
	stateIdx  int
	business  bool
	quitting  bool
	lastWidth int
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func initialModel() model {
	fields := make([]textinput.Model, 4)

	for i := range fields {
		ti := textinput.New()
		ti.Prompt = ""
		ti.Width = 22
		ti.CharLimit = 64
		ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#EAF4FF"))
		ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#CFE8FF"))
		ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#EAF4FF"))
		fields[i] = ti
	}

	fields[firstNameField].Placeholder = ""
	fields[lastNameField].Placeholder = ""
	fields[emailField].Placeholder = ""
	fields[phoneField].Placeholder = ""
	fields[phoneField].Validate = isIntValidator
	fields[firstNameField].Focus()

	return model{
		inputs:   fields,
		focus:    0,
		states:   []string{"AK", "AL", "AR", "AZ", "CA", "CO", "FL", "NY", "TX", "WA"},
		stateIdx: 0,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.quitting {
		return m, tea.Quit
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.lastWidth = msg.Width

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit
		case tea.KeyTab, tea.KeyEnter, tea.KeyDown:
			m.nextField()
			return m, nil
		case tea.KeyShiftTab, tea.KeyUp:
			m.prevField()
			return m, nil
		case tea.KeyLeft:
			if m.focus == stateField {
				if m.stateIdx == 0 {
					m.stateIdx = len(m.states) - 1
				} else {
					m.stateIdx--
				}
				return m, nil
			}
		case tea.KeyRight:
			if m.focus == stateField {
				m.stateIdx = (m.stateIdx + 1) % len(m.states)
				return m, nil
			}
		case tea.KeySpace:
			if m.focus == businessField {
				m.business = !m.business
				return m, nil
			}
		}
	}

	if m.focus >= firstNameField && m.focus <= phoneField {
		var cmd tea.Cmd
		m.inputs[m.focus], cmd = m.inputs[m.focus].Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	rows := []string{
		m.renderInputRow("First Name", firstNameField),
		"",
		m.renderInputRow("Last Name", lastNameField),
		"",
		m.renderInputRow("Email", emailField),
		"",
		m.renderInputRow("Phone", phoneField),
		"",
		m.renderStateRow(),
		"",
		m.renderBusinessRow(),
	}

	return containerStyle.Render(lipgloss.JoinVertical(lipgloss.Left, rows...))
}

func (m model) renderInputRow(label string, idx int) string {
	labelText := labelStyle.Render(label)
	fieldText := m.inputs[idx].View()

	boxStyle := inputStyle
	if m.focus == idx {
		boxStyle = focusedInputStyle
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		labelText,
		boxStyle.Render(fieldText),
	)
}

func (m model) renderStateRow() string {
	labelText := labelStyle.Render("State")
	value := m.states[m.stateIdx]

	style := stateStyle
	if m.focus == stateField {
		style = focusedStateStyle
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, labelText, style.Render(value))
}

func (m model) renderBusinessRow() string {
	labelText := labelStyle.Render("Business")
	mark := " "
	if m.business {
		mark = "x"
	}

	style := checkboxStyle
	if m.focus == businessField {
		style = focusedCheckboxStyle
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, labelText, style.Render(mark))
}

func (m *model) nextField() {
	m.focus = (m.focus + 1) % totalFields
	m.syncFocus()
}

func (m *model) prevField() {
	m.focus--
	if m.focus < 0 {
		m.focus = totalFields - 1
	}
	m.syncFocus()
}

func (m *model) syncFocus() {
	for i := range m.inputs {
		m.inputs[i].Blur()
	}

	if m.focus >= firstNameField && m.focus <= phoneField {
		m.inputs[m.focus].Focus()
	}
}

func (m model) summary() string {
	return fmt.Sprintf(
		"%s %s <%s> %s %s business=%t",
		m.inputs[firstNameField].Value(),
		m.inputs[lastNameField].Value(),
		m.inputs[emailField].Value(),
		m.inputs[phoneField].Value(),
		m.states[m.stateIdx],
		m.business,
	)
}
