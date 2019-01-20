package cmd

import (
	"fmt"

	"github.com/lyraproj/lyra/cmd/lyra/ui"
	"github.com/spf13/cobra"
)

// NewUpgradeCmd returns the register subcommand
func NewUpgradeCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "upgrade",
		Example: "upgrade",
		Short:   "Update server component",
		Long:    "Update server component",
		Run:     runUpgrade,
	}

	cmd.SetHelpTemplate(ui.HelpTemplate)
	cmd.SetUsageTemplate(ui.UsageTemplate)

	return cmd
}

func runUpgrade(cmd *cobra.Command, args []string) {
	fmt.Println("info", "")
	fmt.Println("info", "Upgrade command")
	fmt.Println("info", "")
}
