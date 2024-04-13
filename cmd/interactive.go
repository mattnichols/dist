package cmd

import (
	"dist/lev"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func RunInteractive() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Encountered an error: %v", err)
		os.Exit(1)
	}
}

type styles struct {
	input    lipgloss.Style
	title    lipgloss.Style
	distance lipgloss.Style
}

func defaultStyles() *styles {
	return &styles{
		input:    lipgloss.NewStyle().BorderForeground(lipgloss.Color(205)).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80),
		title:    lipgloss.NewStyle().Align(lipgloss.Center).Padding(2, 0, 0, 0).Width(80),
		distance: lipgloss.NewStyle().Bold(true).Padding(1, 1, 2, 2).Width(80),
	}
}

type model struct {
	a      string
	b      string
	dist   int
	inputs []textinput.Model
	cursor int
	styles *styles
}

func initialModel() *model {
	in := make([]textinput.Model, 2)

	a := textinput.New()
	a.Placeholder = "string 1"
	a.Focus()
	in[0] = a

	b := textinput.New()
	b.Placeholder = "string 2"
	in[1] = b

	return &model{
		inputs: in,
		styles: defaultStyles(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "up", "down", "enter", "tab":
			if m.cursor > 0 {
				m.cursor = 0
			} else {
				m.cursor = 1
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.cursor {
					cmds[i] = m.inputs[i].Focus()
					continue
				}
				m.inputs[i].Blur()
			}

			return m, tea.Batch(cmds...)

		}
	}

	// Update the interface as the user types
	return m, m.updateInterface(msg)
}

func (m *model) updateInterface(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	s := make([]string, 2)
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
		s[i] = m.inputs[i].Value()
	}

	m.dist = lev.Calc(s[0], s[1])

	return tea.Batch(cmds...)
}

func (m model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.styles.title.Render("Levenshtein Distance"),
		m.styles.input.Render(m.inputs[0].View()),
		m.styles.input.Render(m.inputs[1].View()),
		m.styles.distance.Render(fmt.Sprintf("Dist: %d", m.dist)),
	)
}
