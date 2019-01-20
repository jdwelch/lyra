package cmd

import (
	"fmt"

	"github.com/lyraproj/lyra/cmd/lyra/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewInitCmd returns the init subcommand
func NewInitCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:        "init",
		Example:    "init --kube-namespace default",
		Short:      "Install Lyra server component",
		Long:       "Install Lyra server component",
		Run:        runInit,
		SuggestFor: []string{"install"},
	}

	return cmd
}

func runInit(cmd *cobra.Command, args []string) {
	h := viper.GetString("cygnus-host")
	fmt.Println("")
	fmt.Println("Cygnus (the Lyra server components) has been installed into a Kubernetes cluster at " + h + ".")
	fmt.Println("")
	fmt.Println("For more information on using Lyra with Cygnus, see https://github.com/lyraproj/docs/cygnus.md")
}

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
