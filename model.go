package glitter

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	strings.Builder
	wrapped tea.Model
}

// TODO(GIA) Write should also update

func WrapModel(model tea.Model) *Model {
	return &Model{wrapped: model}
}

func (m *Model) Init() tea.Cmd {
	return m.wrapped.Init()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.wrapped, cmd = m.wrapped.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	return m.String() + m.wrapped.View()
}
