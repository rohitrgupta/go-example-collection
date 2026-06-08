package main

import (
	"sort"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	// Keybindings used to enable navigation and filtering when the table is focused.
	Filter key.Binding

	// Keybindings used when setting a filter.
	ClearFilter          key.Binding
	AcceptWhileFiltering key.Binding
}

// DefaultKeyMap returns a default set of keybindings.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Filter: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "filter"),
		),
		ClearFilter: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "clear filter"),
		),
		AcceptWhileFiltering: key.NewBinding(
			key.WithKeys("tab", "shift+tab"),
			key.WithHelp("tab", "Switch to table"),
		),
	}
}

type FilterState int

// Possible filter states.
const (
	Unfiltered    FilterState = iota // no filter set
	Filtering                        // user is actively setting a filter
	FilterApplied                    // a filter is applied and user is not editing filter
)

type FilteredTable struct {
	filterInput    textinput.Model
	previousFilter string
	table          table.Model
	title          string
	columns        []table.Column
	rows           map[int]table.Row
	filterState    FilterState
	focus          bool
	KeyMap         KeyMap
	Help           help.Model
	filterCols     []int
	rowCursor      map[int]int
	cusrorRow      map[int]int
}

func NewFilteredTable(columns []table.Column) FilteredTable {
	ti := textinput.New()
	ti.Prompt = "Filter: "
	return FilteredTable{
		filterInput: ti,
		table: table.New(
			table.WithColumns(columns),
			table.WithFocused(true),
		),
		columns:        columns,
		filterState:    Unfiltered,
		KeyMap:         DefaultKeyMap(),
		title:          lipgloss.NewStyle().Bold(true).Render(""),
		Help:           help.New(),
		previousFilter: "<>",
	}
}

func (m *FilteredTable) SetRows(rows map[int]table.Row) {
	m.rows = rows
	m.FilterRows()
}

func (m *FilteredTable) SetFilterColumns(cols []int) {
	m.filterCols = cols
}

func (m *FilteredTable) FilterRows() {
	if m.previousFilter == m.filterInput.Value() {
		return
	}
	m.previousFilter = m.filterInput.Value()

	cursor := m.table.Cursor()
	// if cursor is in rowCursor, get the corresponding row index in the table and set cursor to that index
	cursorRow := -1

	if c, ok := m.cusrorRow[cursor]; ok {
		cursorRow = c
	}

	m.table.SetRows(nil)
	m.rowCursor = make(map[int]int)
	m.cusrorRow = make(map[int]int)
	keys := m.sortedKeys()
	for _, k := range keys {
		r := m.rows[k]
		// if any of filter columns contain the filter value, include the row in the table
		include := false
		if m.filterInput.Value() == "" {
			include = true
		} else {
			for _, col := range m.filterCols {
				if strings.Contains(r[col], m.filterInput.Value()) {
					include = true
					break
				}
			}
		}
		if include {
			m.table.SetRows(append(m.table.Rows(), r))
			m.rowCursor[k] = len(m.table.Rows()) - 1
			m.cusrorRow[len(m.table.Rows())-1] = k
		}
	}
	// if the selected row is still visible after filtering, keep it selected
	if cursorRow != -1 {
		if c, ok := m.rowCursor[cursorRow]; ok {
			m.table.SetCursor(c)
		} else {
			m.table.SetCursor(0)
		}
	}
}

func (m *FilteredTable) sortedKeys() []int {
	keys := make([]int, 0, len(m.rows))
	for k := range m.rows {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
func (m FilteredTable) Init() tea.Cmd {
	return nil
}

func (m FilteredTable) Update(msg tea.Msg) (FilteredTable, tea.Cmd) {
	// if !m.focus {
	// 	return m, nil
	// }
	var cmds []tea.Cmd

	if m.filterState == Filtering {
		cmds = m.handleFiltering(msg, cmds)
	} else {
		cmds = append(cmds, m.handleBrowsing(msg)...)
	}

	return m, tea.Batch(cmds...)
}

func (m *FilteredTable) handleBrowsing(msg tea.Msg) []tea.Cmd {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Filter):
			m.filterState = Filtering
			m.filterInput.Width = 20
			cmd1 := m.filterInput.Focus()
			cmds = append(cmds, cmd1)
			return cmds
		}
	}
	if len(m.table.Rows()) != 0 {
		var cmd tea.Cmd
		m.table, cmd = m.table.Update(msg)
		cmds = append(cmds, cmd)
		return cmds
	}
	return cmds
}

func (m *FilteredTable) handleFiltering(msg tea.Msg, cmds []tea.Cmd) []tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.ClearFilter):
			m.filterState = Unfiltered
			m.filterInput.SetValue("")
			m.filterInput.Blur()
		case key.Matches(msg, m.KeyMap.AcceptWhileFiltering):
			m.filterState = FilterApplied
			m.filterInput.Blur()
		}
	}
	var cmd tea.Cmd
	m.filterInput, cmd = m.filterInput.Update(msg)
	cmds = append(cmds, cmd)

	// clear table rows and re-populate based on filter
	m.FilterRows()
	return cmds
}

func (m FilteredTable) View() string {

	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.title,
			m.filterInput.View()),
		lipgloss.NewStyle().Height(1).Render(""),
		m.table.View(),
		m.Help.View(m),
	)
}

func (m *FilteredTable) SetStyles(s table.Styles) {
	m.table.SetStyles(s)
}

func (m FilteredTable) Focused() bool {
	return m.focus
}

func (m *FilteredTable) Focus() {
	m.focus = true
	m.table.Focus()
}

func (m *FilteredTable) Blur() {
	m.focus = false
	m.table.Blur()
}

func (m FilteredTable) SelectedRow() table.Row {
	return m.table.SelectedRow()
}

func (m FilteredTable) ShortHelp() []key.Binding {
	if m.filterState == Filtering {
		return []key.Binding{
			m.KeyMap.ClearFilter,
			m.KeyMap.AcceptWhileFiltering,
		}
	}
	kb := []key.Binding{
		m.KeyMap.Filter,
	}

	return kb
}

func (m FilteredTable) FullHelp() [][]key.Binding {
	kb := [][]key.Binding{
		{
			m.KeyMap.Filter,
			m.KeyMap.ClearFilter,
			m.KeyMap.AcceptWhileFiltering,
		},
	}
	return kb
}
