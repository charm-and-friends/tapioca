package glitter

import (
	"bytes"
	"io"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
)

// Adds an io.Writer interface to tea.Program
type Program struct {
	*tea.Program
	buf   bytes.Buffer
	mutex sync.Mutex
}

func (w *Program) Write(p []byte) (int, error) {
	// Buffer everything
	n, err := w.buf.Write(p)
	if err != nil {
		return n, err
	}

	// Write when a new line is found
	w.mutex.Lock()
	defer w.mutex.Unlock()
	for {
		line, err := w.buf.ReadBytes('\n')
		if err == io.EOF {
			// Put back the read data if newline not found.
			w.buf.Write(line)
			return n, nil
		}

		if err != nil {
			return n, err
		}

		// Print everything but the \n
		w.Println(string(line[:len(line)-1]))
	}
}

func NewProgram(model tea.Model, opts ...tea.ProgramOption) *Program {
	return &Program{Program: tea.NewProgram(model, opts...)}
}
