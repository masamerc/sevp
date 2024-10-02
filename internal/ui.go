package internal

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	hexWhite         = "#FFFFFF"
	hexBrightPurpule = "#9370DB"
	hexDarkPurpule   = "#4D1A7F"
	hexBrightGreen   = "#3CCE92"

	defaultWidth = 30
	listHeight   = 15
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2).Bold(true)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color(hexBrightGreen)).Bold(true)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1).Foreground(lipgloss.Color(hexBrightPurpule))
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 1, 4).Foreground(lipgloss.Color(hexBrightGreen))
	plainTextStyle    = lipgloss.NewStyle().Margin(1, 0, 1, 4)

	// Prompt
	promptStyle       = lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color(hexWhite)).Bold(true)
	filerPromptString = "Search: "
	cursorStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color(hexWhite))
)

type item string

func (i item) FilterValue() string { return string(i) }

type itemDelegate struct{}

func (d itemDelegate) Height() int { return 1 }

func (d itemDelegate) Spacing() int { return 0 }

func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := string(i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + s[0])
		}
	}

	fmt.Fprint(w, fn(str))
}

/*
---------- Model Settings ----------
*/

type model struct {
	list     list.Model
	items    []item
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			if m.list.FilterState() != list.Filtering {
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.choice = string(i)
				}
				return m, tea.Quit
			}

		default:
			if !m.list.SettingFilter() && (keypress == "q" || keypress == "esc") {
				m.quitting = true
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		WriteToFile(m.choice)
		return plainTextStyle.Render(fmt.Sprintf("Switched to: %s", m.choice))
	}
	if m.quitting {
		return plainTextStyle.Render("Aborted.")
	}
	return "\n" + m.list.View()
}

/*
---------- App Settings ----------
*/
type App struct {
	items    []string
	teaItems []list.Item
}

func (a *App) Init() {
	teaItems := []list.Item{}

	for _, itemString := range a.items {
		teaItems = append(teaItems, item(itemString))
	}

	a.teaItems = teaItems
}

func (a App) Run() {
	l := list.New(a.teaItems, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "type '/' to search"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	l.FilterInput.Prompt = filerPromptString
	l.FilterInput.PromptStyle = promptStyle
	l.FilterInput.Cursor.Style = cursorStyle

	m := model{list: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func Run(profiles []string) {
	app := App{items: profiles}
	app.Init()
	app.Run()
}
