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

func (program *Program) Write(p []byte) (int, error) {
	// Buffer everything
	n, err := program.buf.Write(p)
	if err != nil {
		return n, err
	}

	// Write when a new line is found
	program.mutex.Lock()
	defer program.mutex.Unlock()
	for {
		line, err := program.buf.ReadBytes('\n')
		if err == io.EOF {
			// Put back the read data if newline not found.
			program.buf.Write(line)
			return n, nil
		}

		if err != nil {
			return n, err
		}

		// Print everything but the \n
		program.Println(string(line[:len(line)-1]))
	}
}

func (program *Program) GoRun(afterRun ...func(tea.Model, error)) *Program {
	go func() {
		m, err := program.Run()
		for _, fn := range afterRun {
			fn(m, err)
		}
	}()
	return program
}

func (program *Program) QuitAndWait() {
	program.Quit()
	program.Wait()
}

func NewProgram(model tea.Model, opts ...tea.ProgramOption) *Program {
	return &Program{Program: tea.NewProgram(model, opts...)}
}
