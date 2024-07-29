package main

import (
	"flag"
	"log"
	"time"

	"github.com/esdandreu/tapioca"
	"github.com/esdandreu/tapioca/progress"
)

var logger *log.Logger = log.Default()

func main() {
	var notapioca bool
	flag.BoolVar(&notapioca, "no-tapioca", false, "Do not use tapioca")
	flag.Parse()

	// Create and start the progress bar
	program := tapioca.NewProgram(progress.New()).GoRun()
	defer program.QuitAndWait()

	if !notapioca {
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
