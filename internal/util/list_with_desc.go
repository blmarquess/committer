package util

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

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
			i, ok := m.list.SelectedItem().(item)
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
		item{title: "✨feat", desc: "Add new feature"},
		item{title: "🐞fix", desc: "Fix a bug"},
		item{title: "🎯test", desc: "Add or update tests"},
		item{title: "♻refactor", desc: "Code changes that neither fixes a bug nor adds a feature"},
		item{title: "🎨style", desc: "Code style changes (whitespace, formatting, etc.)"},
		item{title: "🚧wip", desc: "Work in progress"},
		item{title: "📚docs", desc: "Update documentation"},
		item{title: "📦build", desc: "Changes related to build process"},
		item{title: "♾️ci", desc: "Changes to CI configuration or scripts"},
		item{title: "⚡perf", desc: "Performance improvements"},
		item{title: "↩revert", desc: "Reverts a previous commit"},
		item{title: "🔧chore", desc: "Changes to the build process, auxiliary tools, etc."},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Choice a commit type"

	p := tea.NewProgram(m, tea.WithAltScreen())
	res, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	value := res.View()
	return value
}
