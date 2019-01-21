package cmd

import (
	"errors"

	"github.com/lyraproj/lyra/cmd/lyra/ui"
	m "github.com/lyraproj/lyra/pkg/mock"
	"github.com/spf13/cobra"
)

var (
	yamlOut bool
	jsonOut bool
)

// NewShowCmd returns the show subcommand
func NewShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "show",
		Example: "show",
		Short:   "Enumerate instances of workflows",
		Long:    "Enumerate instances of workflows",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires an argument")
			}
			return nil
		},
		Run:        runShow,
		SuggestFor: []string{"list", "ls", "view"},
	}

	cmd.SetHelpTemplate(ui.HelpTemplate)
	cmd.SetUsageTemplate(ui.UsageTemplate)
	cmd.Flags().BoolVar(&jsonOut, "json", false, "Show output in JSON")
	cmd.Flags().BoolVar(&yamlOut, "yaml", false, "Show output in YAML (default)")

	return cmd
}

func runShow(cmd *cobra.Command, args []string) {
	showResources(args)
}

func showResources(args []string) {
	m.Outline(m.Process, "Try and find the context '"+args[0]+"'")
	ui.DescribeStep("Looking for resources in context '" + args[0] + "'…")
	ui.ProgressBar("Querying running state", 1500, false)
}
