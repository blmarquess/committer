package util

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Item struct {
	title, desc string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.desc }
func (i Item) FilterValue() string { return i.title }

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(Item)
			if ok {
				m.choice = string(i.title)
			}
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		return m.choice
	}
	return docStyle.Render(m.list.View())
}

func GetListWithDesc() string {
	items := []list.Item{
		Item{title: "✨feat", desc: "Add new feature"},
		Item{title: "🐞fix", desc: "Fix a bug"},
		Item{title: "🎯test", desc: "Add or update tests"},
		Item{title: "♻refactor", desc: "Code changes that neither fixes a bug nor adds a feature"},
		Item{title: "🎨style", desc: "Code style changes (whitespace, formatting, etc.)"},
		Item{title: "🚧wip", desc: "Work in progress"},
		Item{title: "📚docs", desc: "Update documentation"},
		Item{title: "📦build", desc: "Changes related to build process"},
		Item{title: "♾️ci", desc: "Changes to CI configuration or scripts"},
		Item{title: "⚡perf", desc: "Performance improvements"},
		Item{title: "↩revert", desc: "Reverts a previous commit"},
		Item{title: "🔧chore", desc: "Changes to the build process, auxiliary tools, etc."},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Choice a commit type"
	m.list.InfiniteScrolling = true

	p := tea.NewProgram(m, tea.WithAltScreen())
	res, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	value := res.View()
	return value
}
