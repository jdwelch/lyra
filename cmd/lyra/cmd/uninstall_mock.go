package cmd

import (
	"fmt"

	"github.com/lyraproj/lyra/cmd/lyra/ui"
	"github.com/spf13/cobra"
)

// NewUninstallCmd removes things from cluster
func NewUninstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "uninstall",
		Example: "uninstall",
		Short:   "Remove server component",
		Long:    "Remove server component",
		Run:     runUninstallCmd,
	}

	cmd.SetHelpTemplate(ui.HelpTemplate)
	cmd.SetUsageTemplate(ui.UsageTemplate)

	return cmd
}

func runUninstallCmd(cmd *cobra.Command, args []string) {
	fmt.Println("uninstall command")
}
