package cmd

import (
	"fmt"

	"github.com/lyraproj/lyra/cmd/lyra/ui"
	"github.com/spf13/cobra"
)

// NewLintCmd returns the init subcommand
func NewLintCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:        "lint",
		Example:    "lint my-workflow/",
		Short:      "Test workflow / wf package",
		Long:       "Test workflow / wf package",
		Run:        runLint,
		SuggestFor: []string{"check"},
	}

	cmd.SetHelpTemplate(ui.HelpTemplate)
	cmd.SetUsageTemplate(ui.UsageTemplate)

	return cmd
}

func runLint(cmd *cobra.Command, args []string) {
	fmt.Println("info", "")
	fmt.Println("info", "Lint command")
	fmt.Println("info", "")
}
