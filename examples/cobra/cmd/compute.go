package cmd

import (
	"io"
	"time"

	"github.com/esdandreu/glitter"
	"github.com/spf13/cobra"
)

// computeCmd represents the progress command
var computeCmd = &cobra.Command{
	Use:   "compute",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create and run a spinner in the background
		program := glitter.NewProgram(glitter.NewSpinner()).GoRun()
		defer program.QuitAndWait() // Quit and wait until printing finishes

		// Set the command output to the program
		defer func(w io.Writer) { cmd.SetOut(w) }(cmd.OutOrStdout())
		cmd.SetOut(program)

		N := 100
		for i := 0; i <= N; i++ {
			time.Sleep(20 * time.Millisecond)
			cmd.Println(i, "/", N)
		}
		cmd.Println("finished computing")
	},
}

func init() {
	rootCmd.AddCommand(computeCmd)
}
