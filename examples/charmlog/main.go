package main

import (
	"flag"
	"time"

	"github.com/charmbracelet/log"
	"github.com/esdandreu/tapioca"
	"github.com/esdandreu/tapioca/spinner"
)

var logger *log.Logger = log.Default()

func main() {
	var notapioca bool
	flag.BoolVar(&notapioca, "no-tapioca", false, "Do not use tapioca")
	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "Print debug logs")
	flag.Parse()

	// Create and run a spinner in the background
	program := tapioca.NewProgram(spinner.New().
		Title("Work in progress"),
	).GoRun()
	defer program.QuitAndWait() // Quit when command ends

	if !notapioca {
		// Use a logger that works together with bubbletea
		defer func(l *log.Logger) { logger = l }(logger)
		logger = log.New(program)
	}

	if verbose {
		logger.SetLevel(log.DebugLevel)
	}
	logger.SetReportTimestamp(false)

	N := 5
	for i := 0; i <= N; i++ {
		if i > 0 {
			logger.Debugf("Task: %d", i)
		}
		time.Sleep(500 * time.Millisecond)
	}
	logger.Info("Work finished!")
}
