package cmd

import (
	"io"
	"time"

	"github.com/esdandreu/glitter"
	"github.com/esdandreu/glitter/spinner"
	"github.com/spf13/cobra"
)

var verbose bool

// connectCmd represents the progress command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create and run a spinner in the background
		program := glitter.NewProgram(spinner.New().Title("Connecting")).GoRun()
		defer program.QuitAndWait() // Quit when command ends

		// Set the command output to the program
		defer func(w io.Writer) { cmd.SetOut(w) }(cmd.OutOrStdout())
		cmd.SetOut(program)

		N := 10
		for i := 0; i <= N; i++ {
			if verbose {
				if i == 0 {
					cmd.Println("First connection attempt")
				} else {
					cmd.Println("Retry connection:", i)
				}
			}
			time.Sleep(500 * time.Millisecond)
		}
		cmd.Println("Connection successful!")
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().BoolVarP(
		&verbose, "verbose", "v", false, "Displays debug logs",
	)
}
