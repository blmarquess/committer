package util

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"

// 	"github.com/charmbracelet/bubbles/list"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )

// const listHeight = 14

// var (
// 	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
// 	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
// 	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
// 	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
// 	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
// 	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
// )

// type item string

// func (i item) FilterValue() string { return "" }

// type itemDelegate struct{}

// func (d itemDelegate) Height() int                             { return 1 }
// func (d itemDelegate) Spacing() int                            { return 0 }
// func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
// func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
// 	i, ok := listItem.(item)
// 	if !ok {
// 		return
// 	}

// 	str := fmt.Sprintf("%d. %s", index+1, i)

// 	fn := itemStyle.Render
// 	if index == m.Index() {
// 		fn = func(s ...string) string {
// 			return selectedItemStyle.Render("> " + strings.Join(s, " "))
// 		}
// 	}

// 	fmt.Fprint(w, fn(str))
// }

// type model struct {
// 	list     list.Model
// 	choice   string
// 	quitting bool
// }

// func (m model) Init() tea.Cmd {
// 	return nil
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		m.list.SetWidth(msg.Width)
// 		return m, nil

// 	case tea.KeyMsg:
// 		switch keypress := msg.String(); keypress {
// 		case "ctrl+c":
// 			m.quitting = true
// 			return m, tea.Quit

// 		case "enter":
// 			i, ok := m.list.SelectedItem().(item)
// 			if ok {
// 				m.choice = string(i)
// 			}
// 			return m, tea.Quit
// 		}
// 	}

// 	var cmd tea.Cmd
// 	m.list, cmd = m.list.Update(msg)
// 	return m, cmd
// }

// func (m model) View() string {
// 	if m.choice != "" {
// 		return m.choice
// 	}
// 	if m.quitting {
// 		return quitTextStyle.Render("Not hungry? That’s cool.")
// 	}
// 	return "\n" + m.list.View()
// }

// type ListOptions struct {
// 	Items []string
// 	Title string
// }

// func CreateList(options ListOptions) (string, error) {
// 	var items []list.Item
// 	for _, i := range options.Items {
// 		items = append(items, item(i))
// 	}

// 	const defaultWidth = 20

// 	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
// 	l.Title = options.Title
// 	l.SetShowStatusBar(false)
// 	l.SetFilteringEnabled(false)
// 	l.Styles.Title = titleStyle
// 	l.Styles.PaginationStyle = paginationStyle
// 	l.Styles.HelpStyle = helpStyle

// 	m := model{list: l}

// 	res, err := tea.NewProgram(m).Run()
// 	if err != nil {
// 		fmt.Println("Error running program:", err)
// 		os.Exit(1)
// 	}
// 	res.View()
// 	fmt.Println(l.SelectedItem().FilterValue())
// 	fmt.Println(m.list.SelectedItem().FilterValue())
// 	la := res.View()
// 	return la, err
// }

// options := modules.ListOptions{
// 	Items: []modules.Item{
// 		modules.Item{Title: "feat", Desc: "nova funcionalidade"},
// 		modules.Item{Title: "fix", Desc: "descrição"},
// 		modules.Item{Title: "docs", Desc: "descrição"},
// 		modules.Item{Title: "style", Desc: "descrição"},
// 		modules.Item{Title: "refactor", Desc: "descrição"},
// 		modules.Item{Title: "test", Desc: "descrição"},
// 		modules.Item{Title: "chore", Desc: "descrição"},
// 		modules.Item{Title: "build", Desc: "descrição"},
// 		modules.Item{Title: "ci", Desc: "descrição"},
// 		modules.Item{Title: "perf", Desc: "descrição"},
// 		modules.Item{Title: "revert", Desc: "descrição"},
// 	},
// 	Title: "Qual o tipo de commit?",
// }
// options := modules.ListOptions{
// 	Items: []string{"feat", "fix", "docs", "style", "refactor", "test", "chore", "build", "ci", "perf", "revert"},
// 	Title: "Qual o tipo de commit?",
// }
// commitType, err := modules.CreateList(options)

// if err != nil {
// 	fmt.Printf("Error: %s\n", err)
// 	os.Exit(1)
// }
