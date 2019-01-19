package cmd

import (
	"fmt"

	t "github.com/lyraproj/lyra/pkg/strings"
	"github.com/spf13/cobra"
)

// NewLintCmd returns the init subcommand
func NewLintCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:        "lint",
		Example:    "lint",
		Short:      "lint",
		Long:       "lint",
		Run:        runLint,
		SuggestFor: []string{"check"},
	}

	cmd.SetHelpTemplate(t.HelpTemplate)
	cmd.SetUsageTemplate(t.UsageTemplate)

	return cmd
}

func runLint(cmd *cobra.Command, args []string) {
	fmt.Println("info", "")
	fmt.Println("info", "Lint command")
	fmt.Println("info", "")
}
