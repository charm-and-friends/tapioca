package glitter

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Handles common CLI signals like Ctrl+C for quitting.
func HandleMessage(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return tea.Quit
		}
	}
	return nil
}

type ModelWrapper struct {
	tea.Model
}

func (m *ModelWrapper) Init() tea.Cmd {
	return m.Model.Init()
}

func (m *ModelWrapper) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := HandleMessage(msg)
	if cmd != nil {
		return m, cmd
	}
	m.Model, cmd = m.Model.Update(msg)
	return m, cmd
}

func (m *ModelWrapper) View() string {
	return m.Model.View()
}

func WrapModel(model tea.Model) *ModelWrapper {
	return &ModelWrapper{model}
}
