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
		Example:    "reg",
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
	fmt.Println("info", "")
	fmt.Println("info", "Register command")
	fmt.Println("info", "")
}

// NewUnregisterCmd returns the unregister subcommand
func NewUnregisterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:        "unregister",
		Example:    "unreg",
		Short:      "Remove workflow from a cluster",
		Long:       "Remove workflow from a cluster",
		Run:        runRegister,
		SuggestFor: []string{"undeploy"},
	}

	cmd.SetHelpTemplate(ui.HelpTemplate)
	cmd.SetUsageTemplate(ui.UsageTemplate)

	return cmd
}

func runUnRegister(cmd *cobra.Command, args []string) {
	fmt.Println("info", "")
	fmt.Println("info", "Unregister command")
	fmt.Println("info", "")
}
