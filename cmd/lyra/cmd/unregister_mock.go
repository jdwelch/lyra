package cmd

import (
	"fmt"

	"github.com/lyraproj/lyra/cmd/lyra/ui"
	"github.com/spf13/cobra"
)

// NewUnregisterCmd returns the unregister subcommand
func NewUnregisterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:        "unregister",
		Example:    "unregister",
		Short:      "Remove workflow from a cluster",
		Long:       "Remove workflow from a cluster",
		Run:        runUnRegister,
		SuggestFor: []string{"undeploy"},
	}

	cmd.SetHelpTemplate(ui.HelpTemplate)
	cmd.SetUsageTemplate(ui.UsageTemplate)

	return cmd
}

func runUnRegister(cmd *cobra.Command, args []string) {
	fmt.Println("info", "Unregister command")
}
