package spinner

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esdandreu/glitter"
)

// Wraps bubbles/spinner.Spinner in a way that it implements tea.Spinner interface.
type Spinner struct {
	spinner.Model
}

func (m Spinner) Init() tea.Cmd {
	return m.Model.Tick
}

func (m Spinner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Model, cmd = m.Model.Update(msg)
	return m, cmd
}

func (m Spinner) View() string {
	return m.Model.View()
}

// Creates a new Spinner tea.Model.
func New(opts ...spinner.Option) tea.Model {
	return glitter.WrapModel(&Spinner{spinner.New(opts...)})
}
