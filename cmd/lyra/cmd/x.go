package cmd

import (
	m "github.com/lyraproj/lyra/pkg/mock"
	"github.com/spf13/cobra"
)

// NewExperimentalCmd returns the x subcommand
// NB: It's hidden from the help menu, on purpose. Sneaky.
func NewExperimentalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "x",
		Hidden: true,
		Short:  "Parent of experimental subcommands",
		Run:    runX,
	}

	return cmd
}

func runX(cmd *cobra.Command, args []string) {
	m.DynamicFeedbackOnProgress()
	m.TableProgress()
}
