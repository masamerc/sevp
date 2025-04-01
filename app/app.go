package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	items    []string
	teaItems []list.Item
	target   string
}

func NewApp(items []string, targetVar string) *App {
	teaItems := make([]list.Item, len(items))
	for i, itemString := range items {
		teaItems[i] = Item(itemString)
	}
	return &App{
		items:    items,
		teaItems: teaItems,
		target:   targetVar,
	}
}

func (a *App) Run() error {
	listStyles := NewStyles().List
	renderStyles := NewStyles().Rendering

	l := list.New(a.teaItems, NewItemDelegate(), DefaultWidth, ListHeight)

	// title setting
	title := fmt.Sprintf("[%s]\n\ntype '/' to search", renderStyles.TargetType.Render(a.target))
	l.Title = title

	// general settings
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)

	// list styling
	l.Styles.Title = listStyles.Styles.Title
	l.Styles.HelpStyle = listStyles.Styles.HelpStyle
	l.FilterInput.Prompt = "Search: "
	l.FilterInput.PromptStyle = listStyles.Styles.FilterPrompt
	l.FilterInput.Cursor.Style = listStyles.Styles.FilterPrompt
	l.Styles.FilterCursor = listStyles.Styles.FilterCursor
	l.Styles.PaginationStyle = listStyles.Styles.PaginationStyle

	m := NewModel(l, a.target)

	_, err := tea.NewProgram(m).Run()
	return err
}
