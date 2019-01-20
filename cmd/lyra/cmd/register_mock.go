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
		Short:      "reg",
		Long:       "reg",
		Run:        runRegister,
		SuggestFor: []string{"deploy"},
	}

	cmd.SetHelpTemplate(t.HelpTemplate)
	cmd.SetUsageTemplate(t.UsageTemplate)

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
		Short:      "unreg",
		Long:       "unreg",
		Run:        runRegister,
		SuggestFor: []string{"undeploy"},
	}

	cmd.SetHelpTemplate(t.HelpTemplate)
	cmd.SetUsageTemplate(t.UsageTemplate)

	return cmd
}

func runUnRegister(cmd *cobra.Command, args []string) {
	fmt.Println("info", "")
	fmt.Println("info", "Unregister command")
	fmt.Println("info", "")
}
