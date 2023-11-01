package main

import (
	"flag"
	"log"
	"time"

	"github.com/esdandreu/glitter"
)

var logger *log.Logger = log.Default()

func main() {
	var noGlitter bool
	flag.BoolVar(&noGlitter, "no-glitter", false, "Do not use glitter")
	flag.Parse()

	// Create and start the progress bar
	program := glitter.NewProgram(glitter.NewProgress()).GoRun()
	defer program.QuitAndWait()

	if !noGlitter {
		// Use a logger that works together with bubbletea
		defer func(l *log.Logger) { logger = l }(logger)
		logger = log.New(program, "", log.LstdFlags)
	}

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
