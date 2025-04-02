package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/masamerc/sevp/internal"
)

type Model struct {
	list     list.Model
	choice   string
	quitting bool
	target   string
}

func NewModel(l list.Model, target string) Model {
	return Model{list: l, target: target}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			if m.list.FilterState() != list.Filtering {
				i, ok := m.list.SelectedItem().(Item)
				if ok {
					m.choice = string(i)
				}
				return m, tea.Quit
			}
		default:
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

func (m Model) View() string {
	renderStyles := NewStyleSet().Rendering
	if m.choice != "" {
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
		return renderStyles.QuitText.Render("Aborted.")
	}
	return "\n" + m.list.View()
}
