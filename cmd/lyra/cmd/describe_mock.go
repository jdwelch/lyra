package cmd

import (
	"fmt"

	"github.com/lyraproj/lyra/cmd/lyra/ui"
	"github.com/spf13/cobra"
)

// NewDescribeCmd returns the register subcommand
func NewDescribeCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "describe",
		Example: "describe",
		Short:   "Enumerate details of an instance of a workflow",
		Long:    "Enumerate details of an instance of a workflow",
		Run:     runDescribe,
	}

	cmd.SetHelpTemplate(ui.HelpTemplate)
	cmd.SetUsageTemplate(ui.UsageTemplate)

	return cmd
}

func runDescribe(cmd *cobra.Command, args []string) {
	fmt.Println("info", "")
	fmt.Println("info", "Describe command")
	fmt.Println("info", "")
}
