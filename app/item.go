package app

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Item represents a single item in the list
type Item string

// FilterValue returns the string representation of the item
func (i Item) FilterValue() string { return string(i) }

// ItemDelegate is a custom delegate for rendering items in the list
type ItemDelegate struct{}

// NewItemDelegate creates a new empty instance of ItemDelegate
func NewItemDelegate() ItemDelegate {
	return ItemDelegate{}
}

// Height defines the height of the item in the list
func (d ItemDelegate) Height() int { return 1 }

// Spacing defines the spacing between items in the list
func (d ItemDelegate) Spacing() int { return 0 }

// Update is a no-op for the item delegate
func (d ItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

// Render is the main rendering function for the item delegate.
// It determines how the item should be displayed based on its index.
func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	renderStyles := NewStyleSet().Rendering
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := string(i)

	// default render function / style for each item
	fn := renderStyles.Item.Render

	// if the item is selected, apply a different style and format
	if index == m.Index() {
		fn = func(strs ...string) string {
			return renderStyles.SelectedItem.Render("> " + strs[0])
		}
	}

	// render the item with the selected style
	fmt.Fprint(w, fn(str))
}
