package glitter_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/esdandreu/glitter"
	"github.com/esdandreu/glitter/spinner"
	"github.com/stretchr/testify/assert"
)

func ExampleProgram() {
	program := glitter.NewProgram(spinner.New()).GoRun()
	defer program.QuitAndWait()

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Fprintln(program, i, "milliseconds")
	}
}

func TestProgram(t *testing.T) {
	var buffer bytes.Buffer
	program := glitter.NewProgram(spinner.New(), tea.WithOutput(&buffer))
	// Start the program and log concurrently
	n := 100
	go func() {
		for i := 0; i < n; i++ {
			fmt.Fprintln(program, "DEBUG", i)
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
