package cmd

import (
	"errors"

	"github.com/lyraproj/lyra/pkg/logger"
	t "github.com/lyraproj/lyra/pkg/strings"

	"github.com/spf13/cobra"
)

// NewApplyEnvCmd returns a subcommand of a subcommand
func NewApplyEnvCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "env",
		Example: "  lyra apply env development",
		Short:   "I am a placeholder for a subcommand of a subcommand",
		Long:    "I am a placeholder for a subcommand of a subcommand",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires an argument")
			}
			return nil
		},
		Run:    runApplyEnv,
		Hidden: true,
	}

	cmd.SetHelpTemplate(t.HelpTemplate)
	cmd.SetUsageTemplate(t.UsageTemplate)

	return cmd
}

func runApplyEnv(cmd *cobra.Command, args []string) {
	log := logger.Get()
	log.Error("NOT IMPLEMENTED")
}
