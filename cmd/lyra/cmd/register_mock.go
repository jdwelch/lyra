package cmd

import (
	"fmt"

	t "github.com/lyraproj/lyra/pkg/strings"
	"github.com/spf13/cobra"
)

// NewRegisterCmd returns the init subcommand
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
