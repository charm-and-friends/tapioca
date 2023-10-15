package cmd

import (
	"io"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/esdandreu/sunkentea"
	"github.com/spf13/cobra"
)

// progressCmd represents the progress command
var progressCmd = &cobra.Command{
	Use:   "progress",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create progress bar and set the command output
		defer func(w io.Writer) { cmd.SetOut(w) }(cmd.OutOrStdout())
		pbar, w := sunkentea.NewProgram(&model{
			progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C")),
		})
		cmd.SetOut(w)

		// Start the progress bar
		go func() {
			if _, err := pbar.Run(); err != nil {
				cmd.PrintErr("Oh no!", err)
			}
		}()
		defer pbar.Quit()

		for i := 0; i <= 100; i++ {
			time.Sleep(20 * time.Millisecond)
			cmd.Println("DEBUG", i)
			pbar.Send(float64(i) / 100)
		}
		cmd.Println("progress called")
	},
}

type model struct {
	progress progress.Model
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

func init() {
	rootCmd.AddCommand(progressCmd)
}
