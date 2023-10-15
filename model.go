package sunkentea

import (
	tea "github.com/charmbracelet/bubbletea"
)

type WrittenMsg int

type SunkenModel struct {
	model tea.Model
	// strings.Builder
	buffer string
}

// func (m SunkenModel) Write(p []byte) (int, error) {
// 	n, err := m.Builder.Write(p)
// 	// fmt.Print(m.Builder.String())
// 	m.Update(WrittenMsg(n))
// 	return n, err
// }

func NewSunkenModel(model tea.Model) SunkenModel {
	return SunkenModel{model: model}
}

func (m SunkenModel) Init() tea.Cmd {
	return m.model.Init()
}

func (m SunkenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case string:
		m.buffer += msg
		return m, nil
	// case WrittenMsg:
	// 	f, err := os.OpenFile("log.log", os.O_APPEND, os.ModeAppend)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer f.Close()
	// 	fmt.Fprintln(f, "Update!", m.Builder.String())
	// 	return m, nil
	default:
		var cmd tea.Cmd
		m.model, cmd = m.model.Update(msg)
		return m, cmd
	}
}

func (m SunkenModel) View() string {
	return m.buffer + m.model.View()
}
