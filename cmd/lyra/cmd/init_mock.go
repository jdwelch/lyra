package cmd

import (
	"fmt"

	t "github.com/lyraproj/lyra/pkg/strings"
	"github.com/spf13/cobra"
)

// NewInitCmd returns the init subcommand
func NewInitCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:        "init",
		Example:    "init",
		Short:      "init",
		Long:       "init",
		Run:        runInit,
		SuggestFor: []string{"install"},
	}

	cmd.SetHelpTemplate(t.HelpTemplate)
	cmd.SetUsageTemplate(t.UsageTemplate)

	return cmd
}

func runInit(cmd *cobra.Command, args []string) {
	fmt.Println("info", "")
	fmt.Println("info", "Now installing server component(s) into kube cluster")
	fmt.Println("info", "")
}
