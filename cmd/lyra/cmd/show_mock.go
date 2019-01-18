package cmd

import (
	"errors"
	"fmt"

	"github.com/lyraproj/lyra/cmd/lyra/ui"
	m "github.com/lyraproj/lyra/pkg/mock"
	t "github.com/lyraproj/lyra/pkg/strings"
	"github.com/spf13/cobra"
)

var (
	yamlOut bool
	jsonOut bool
)

// NewShowCmd returns the show subcommand
func NewShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     t.ShowCmdUsage,
		Example: t.ShowCmdExample,
		Short:   t.ShowCmdShortDesc,
		Long:    t.ShowCmdLongDesc,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires an argument")
			}
			return nil
		},
		Run:        runShow,
		SuggestFor: []string{"view", "describe"},
	}

	cmd.SetHelpTemplate(t.HelpTemplate)
	cmd.SetUsageTemplate(t.UsageTemplate)
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
	if jsonOut {
		fmt.Println(jsonOutput)
	} else {
		fmt.Println(yamlOutput)
	}
}

const (
	yamlOutput = `meta:
  cloudcontext:
    key1: val1
    key2: val2
resources:
  namespace::type1:
    'title1':
      key1: value1
      key2: value2
  namespace::type2:
    'title2':
      key1: value1
      key2: value2
      key3: value3
      key4: value4`

	jsonOutput = `{
  "meta": {
    "cloudcontext": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "resources": {
    "namespace::type1": {
      "title1": {
      "key1": "value1"
      }
    },
    "namespace::type2": {
      "title2": {
        "key1": "value1",
        "key2": "value2",
        "key3": "value3",
        "key4": "value4"
      }
    }
  }
}`
)
