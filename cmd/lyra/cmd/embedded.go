package cmd

import (
	"os"

	"github.com/lyraproj/lyra/cmd/goplugin-aws/aws"
	"github.com/lyraproj/lyra/cmd/goplugin-example/example"
	"github.com/lyraproj/lyra/cmd/goplugin-identity/identity"
	tfaws "github.com/lyraproj/lyra/cmd/goplugin-tf-aws/handler"
	tfazurerm "github.com/lyraproj/lyra/cmd/goplugin-tf-azurerm/handler"
	tfgithub "github.com/lyraproj/lyra/cmd/goplugin-tf-github/handler"
	tfgoogle "github.com/lyraproj/lyra/cmd/goplugin-tf-google/handler"
	tfkubernetes "github.com/lyraproj/lyra/cmd/goplugin-tf-kubernetes/handler"
	"github.com/lyraproj/lyra/pkg/logger"
	"github.com/lyraproj/puppet-workflow/puppet"
	"github.com/spf13/cobra"
)

// EmbeddedPluginCmd runs embedded plugins
func EmbeddedPluginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "plugin",
		Hidden: true,
		Run:    startPlugin,
		Args:   cobra.ExactArgs(1),
	}

	cmd.SetHelpTemplate(cmd.HelpTemplate())

	return cmd
}

func startPlugin(cmd *cobra.Command, args []string) {
	name := args[0]
	switch name {
	case "example":
		example.Start()
	case "identity":
		identity.Start("identity.db")
	case "aws":
		aws.Start()
	case "tfaws":
		tfaws.Start()
	case "tfazurerm":
		tfazurerm.Start()
	case "tfgithub":
		tfgithub.Start()
	case "tfgoogle":
		tfgoogle.Start()
	case "tfkubernetes":
		tfkubernetes.Start()
	case "puppet":
		puppet.Start(`Puppet`)
	default:
		logger.Get().Error("Unknown embedded plugin", "name", name)
		os.Exit(1)
	}
}
