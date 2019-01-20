package cmd

import (
	"fmt"

	"github.com/lyraproj/lyra/cmd/lyra/ui"
	"github.com/spf13/cobra"
)

// NewStatusCmd returns the register subcommand
func NewStatusCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "status",
		Example: "status",
		Short:   "Show the state of server component",
		Long:    "Show the state of server component",
		Run:     runStatus,
	}

	cmd.SetHelpTemplate(ui.HelpTemplate)
	cmd.SetUsageTemplate(ui.UsageTemplate)

	return cmd
}

func runStatus(cmd *cobra.Command, args []string) {
	fmt.Println("info", "")
	fmt.Println("info", "Status command")
	fmt.Println("info", "")
}
