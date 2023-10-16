package glitter_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esdandreu/glitter"
	"github.com/stretchr/testify/assert"
)

func TestProgram(t *testing.T) {
	s := model{spinner.New()}
	var buffer bytes.Buffer
	program, writer := glitter.NewProgram(s, tea.WithOutput(&buffer))
	// Start the program and log concurrently
	n := 100
	go func() {
		for i := 0; i < n; i++ {
			time.Sleep(1 * time.Millisecond)
			fmt.Fprintln(writer, "DEBUG", i)
		}
		program.Quit()
	}()
	if _, err := program.Run(); err != nil {
		t.Fatal("program.Run failed:", err)
	}

	// Assert output is expected
	t.Log(buffer.String())
	lines := strings.Split(buffer.String(), "\n")
	assert.Equal(t, n+1, len(lines))
}

type model struct {
	spinner.Model
}

func (m model) Init() tea.Cmd {
	return m.Model.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.Model, cmd = m.Model.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.Model.View()
}
