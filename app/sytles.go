package app

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const (
	HexWhite        = "#FFFFFF"
	HexBrightPurple = "#B198E5"
	HexBrightGreen  = "#3CCE92"

	DefaultWidth = 30
	ListHeight   = 15
)

// RenderingStyles defines styles used for rendering individual items or custom text.
type RenderingStyles struct {
	Item           lipgloss.Style
	SelectedItem   lipgloss.Style
	SelectedResult lipgloss.Style
	QuitText       lipgloss.Style
	PlainText      lipgloss.Style
	TargetType     lipgloss.Style
}

// ListStyles defines styles applied directly to the list.Model.
type ListStyles struct {
	Styles list.Styles
}

// Styles bundles both RenderingStyles and ListStyles for easy access.
type Styles struct {
	Rendering RenderingStyles
	List      ListStyles
}

// NewStyles initializes and returns the full styles set.
func NewStyles() *Styles {
	return &Styles{
		Rendering: RenderingStyles{
			Item:           lipgloss.NewStyle().PaddingLeft(4),
			SelectedItem:   lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color(HexBrightGreen)).Bold(true),
			SelectedResult: lipgloss.NewStyle().Foreground(lipgloss.Color(HexBrightGreen)).Bold(true),
			QuitText:       lipgloss.NewStyle().Margin(1, 0, 1, 4),
			PlainText:      lipgloss.NewStyle().Margin(1, 0, 1, 4),
			TargetType:     lipgloss.NewStyle().Foreground(lipgloss.Color(HexBrightPurple)).Bold(true),
		},
		List: ListStyles{
			Styles: list.Styles{
				Title:           lipgloss.NewStyle().MarginLeft(2).Bold(true),
				PaginationStyle: lipgloss.NewStyle().PaddingLeft(4),
				HelpStyle:       lipgloss.NewStyle().PaddingLeft(4).PaddingBottom(1).Foreground(lipgloss.Color(HexBrightPurple)),
				FilterPrompt:    lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color(HexWhite)).Bold(true),
				FilterCursor:    lipgloss.NewStyle().Foreground(lipgloss.Color(HexWhite)),
			},
		},
	}
}
