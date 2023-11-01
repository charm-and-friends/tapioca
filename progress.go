package glitter

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type Progress struct {
	progress.Model
}

func (m Progress) Init() tea.Cmd {
	return nil
}

func (m Progress) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case float64:
		return m, m.Model.SetPercent(msg)
	}
	// ! Workaround: progress.Model.Update() doesn't return a progress.Model
	updatedModel, cmd := m.Model.Update(msg)
	if progressModel, ok := updatedModel.(progress.Model); ok {
		m.Model = progressModel
	}
	return m, cmd
}

func (m Progress) View() string {
	return m.Model.View()
}

// Creates a new Progress tea.Model.
func NewProgress(opts ...progress.Option) tea.Model {
	return WrapModel(&Progress{progress.New(opts...)})
}
