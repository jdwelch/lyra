package cmd

import (
	"fmt"

	"github.com/lyraproj/lyra/cmd/lyra/ui"
	"github.com/spf13/cobra"
)

// NewRegisterCmd returns the register subcommand
func NewRegisterCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:        "register",
		Example:    "register",
		Short:      "Make workflow available in a cluster",
		Long:       "Make workflow available in a cluster",
		Run:        runRegister,
		SuggestFor: []string{"deploy"},
	}

	cmd.SetHelpTemplate(ui.HelpTemplate)
	cmd.SetUsageTemplate(ui.UsageTemplate)

	return cmd
}

func runRegister(cmd *cobra.Command, args []string) {
	fmt.Println("info", "Register command")
}
