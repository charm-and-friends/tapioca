package main

import (
	"flag"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esdandreu/glitter"
)

var logger *log.Logger = log.Default()

func main() {
	var without bool
	flag.BoolVar(&without, "without", false, "Do not use glitter")
	flag.Parse()

	program, writer := glitter.NewProgram(model{progress.New()})

	if !without {
		// Use a logger that works together with bubbletea
		defer func(l *log.Logger) { logger = l }(logger)
		logger = log.New(writer, "", log.LstdFlags)
	}

	// Start the progress bar
	go func() {
		if _, err := program.Run(); err != nil {
			logger.Fatal("program.Run failed:", err)
		}
	}()
	defer func() {
		program.Quit()
		program.Wait()
	}()

	// Do work, log and increase progress bar
	for i := 0; i <= 100; i++ {
		time.Sleep(10 * time.Millisecond)
		logger.Println("Started ", i)
		time.Sleep(10 * time.Millisecond)
		logger.Println("Finished", i)
		program.Send(float64(i) / 100)
	}
	logger.Println("Finished everything!")
}

type model struct {
	progress.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		logger.Println("Received Key:", msg.String())
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	case float64:
		return m, m.Model.SetPercent(msg)
	}
	updatedModel, cmd := m.Model.Update(msg)
	if progressModel, ok := updatedModel.(progress.Model); ok {
		m.Model = progressModel
	}
	return m, cmd
}

func (m model) View() string {
	return m.Model.View()
}
