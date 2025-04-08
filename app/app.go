package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// App is the main application struct that holds the items to be displayed
type App struct {
	items    []string
	teaItems []list.Item
	target   string
}

// NewApp initializes a new App instance with the provided items and target variable
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

// Run starts the Bubble Tea program with the items and target variable
// set in the App struct
func (a *App) Run() error {
	listStyles := NewStyleSet().List
	renderStyles := NewStyleSet().Rendering

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
