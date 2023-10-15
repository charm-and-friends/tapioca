package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esdandreu/sunkentea"
)

func main() {
	var wrap bool
	flag.BoolVar(&wrap, "wrap", true, "Use sunkentea to wrap the progress bar")
	flag.Parse()

	var model tea.Model = &Model{
		progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C")),
	}
	var output io.Writer = os.Stdout
	if wrap {
		wrapped := sunkentea.WrapModel(model)
		model = wrapped
		output = wrapped
	}

	program := tea.NewProgram(model)

	go func() {
		for i := 0; i <= 100; i++ {
			time.Sleep(20 * time.Millisecond)
			fmt.Fprintln(output, "DEBUG", i)
			program.Send(float64(i) / 100)
		}
		program.Quit()
	}()

	// This is a blocking call
	if _, err := program.Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}

type Model struct {
	progress progress.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case float64:
		m.progress.SetPercent(msg)
		return m, nil
	default:
		return m, nil
	}
}

func (m Model) View() string {
	return m.progress.ViewAs(m.progress.Percent())
}
