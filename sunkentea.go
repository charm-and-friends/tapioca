package sunkentea

import (
	"io"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

func NewProgram(
	w io.Writer, model tea.Model, opts ...tea.ProgramOption,
) (*tea.Program, io.Writer) {
	output := termenv.NewOutput(w, termenv.WithColorCache(true))
	// TODO(GIA) Wrap model
	p := tea.NewProgram(
		model, append(opts, tea.WithOutput(output), tea.WithoutRenderer())...,
	)
	return p, w
}
