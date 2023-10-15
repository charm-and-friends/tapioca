package sunkentea

import (
	"io"

	tea "github.com/charmbracelet/bubbletea"
)

func NewProgram(model tea.Model, opts ...tea.ProgramOption) (*tea.Program, io.Writer) {
	w := WrapModel(model)
	return tea.NewProgram(w, opts...), w
}
