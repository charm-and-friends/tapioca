package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esdandreu/sunkentea"
)

func main() {

	sunkenmodel := sunkentea.WrapModel(&model{
		progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C")),
		true,
	})

	program := tea.NewProgram(sunkenmodel)

	go func() {
		for i := 0; i <= 100; i++ {
			time.Sleep(20 * time.Millisecond)
			fmt.Fprintln(sunkenmodel, "DEBUG", i)
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

type model struct {
	progress progress.Model
	show     bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() string {
	return m.progress.ViewAs(m.progress.Percent())
}
