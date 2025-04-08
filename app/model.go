package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/masamerc/sevp/internal"
)

// Model controls the state of the TUI application
type Model struct {
	list     list.Model
	choice   string
	quitting bool
	target   string
}

// NewModel creates a new instance of the Model with the provided list and target variable
func NewModel(l list.Model, target string) Model {
	return Model{list: l, target: target}
}

// Init is a no-op for the model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update processes messages and updates the model state
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "ctrl+c":
			// CTRL+C always quits the application
			m.quitting = true
			return m, tea.Quit
		case "enter":
			if m.list.FilterState() != list.Filtering {
				// If not in filtering mode, and users presses enter, we want to select the item
				i, ok := m.list.SelectedItem().(Item)
				if ok {
					m.choice = string(i)
				}
				return m, tea.Quit
			}
		default:
			// If not in filtering mode, and users presses 'q' or 'esc', we want to quit
			if !m.list.SettingFilter() && (key == "q" || key == "esc") {
				m.quitting = true
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View defines what the model should display and
// how to terminate the application based on user actions.
func (m Model) View() string {
	renderStyles := NewStyleSet().Rendering

	if m.choice != "" {
		// if users made a selection, we want to write the selected item to the target file
		// and quit the application
		err := internal.WriteToFile(m.choice, m.target)
		if err != nil {
			return renderStyles.QuitText.Render("Error writing to file: " + err.Error())
		}
		return renderStyles.PlainText.Render(
			fmt.Sprintf(
				"%s selected: %s",
				renderStyles.TargetType.Render(m.target),
				renderStyles.SelectedResult.Render(m.choice),
			),
		)
	}

	if m.quitting {
		// we want to quit the application without making a selection
		return renderStyles.QuitText.Render("Aborted.")
	}

	// no selection made, and not quitting -> show the list
	return "\n" + m.list.View()
}
