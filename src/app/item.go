package app

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Item string

func (i Item) FilterValue() string { return string(i) }

type ItemDelegate struct{}

func NewItemDelegate() ItemDelegate {
	return ItemDelegate{}
}

func (d ItemDelegate) Height() int                               { return 1 }
func (d ItemDelegate) Spacing() int                              { return 0 }
func (d ItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	renderStyles := NewStyles().Rendering
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := string(i)
	fn := renderStyles.Item.Render
	if index == m.Index() {
		fn = func(strs ...string) string {
			return renderStyles.SelectedItem.Render("> " + strs[0])
		}
	}

	fmt.Fprint(w, fn(str))
}
