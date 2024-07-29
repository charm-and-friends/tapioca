package spinner

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/esdandreu/tapioca"
)

// Wraps bubbles/spinner.Spinner in a way that it implements tea.Spinner
// interface. It adds a title with the same implementation as huh/spinner.
type Spinner struct {
	spinner.Model
	title      string
	titleStyle lipgloss.Style
}

// Title sets the title of the spinner.
func (m *Spinner) Title(title string) *Spinner {
	m.title = title
	return m
}

// Style sets the style of the spinner.
func (m *Spinner) Style(style lipgloss.Style) *Spinner {
	m.Model.Style = style
	return m
}

// TitleStyle sets the title style of the spinner.
func (m *Spinner) TitleStyle(style lipgloss.Style) *Spinner {
	m.titleStyle = style
	return m
}

func (m Spinner) Init() tea.Cmd {
	return m.Model.Tick
}

func (m Spinner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := tapioca.HandleMessage(msg)
	if cmd != nil {
		return m, cmd
	}
	m.Model, cmd = m.Model.Update(msg)
	return m, cmd
}

func (m Spinner) View() string {
	var title string
	if m.title != "" {
		title = " " + m.titleStyle.Render(m.title)
	}
	return m.Model.View() + title
}

// Creates a new Spinner tea.Model.
func New(opts ...spinner.Option) *Spinner {
	return &Spinner{
		Model:      spinner.New(opts...),
		title:      "",
		titleStyle: lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}),
	}
}
