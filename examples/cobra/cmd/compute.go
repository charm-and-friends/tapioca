package cmd

import (
	"io"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esdandreu/sunkentea"
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
		// Create progress bar and set the command output
		defer func(w io.Writer) { cmd.SetOut(w) }(cmd.OutOrStdout())
		program, w := sunkentea.NewProgram(&model{spinner.New()})
		cmd.SetOut(w)

		// Start the progress bar
		go func() {
			if _, err := program.Run(); err != nil {
				cmd.PrintErr("Oh no!", err)
			}
		}()
		defer func() {
			program.Quit() // Sends a Quit signal
			program.Wait() // Waits until everything has been printed
		}()

		for i := 0; i <= 100; i++ {
			time.Sleep(20 * time.Millisecond)
			cmd.Println("Computed:", i)
		}
		cmd.Println("finished computing")
	},
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

func init() {
	rootCmd.AddCommand(computeCmd)
}
