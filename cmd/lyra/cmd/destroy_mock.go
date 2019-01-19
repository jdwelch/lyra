package cmd

import (
	"errors"
	"fmt"

	ui "github.com/lyraproj/lyra/cmd/lyra/ui"
	m "github.com/lyraproj/lyra/pkg/mock"
	t "github.com/lyraproj/lyra/pkg/strings"

	"github.com/spf13/cobra"
)

// NewDestroyCmd returns the destroy subcommand
func NewDestroyCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     t.DestroyCmdUsage,
		Example: t.DestroyCmdExample,
		Short:   t.DestroyCmdShortDesc,
		Long:    t.DestroyCmdLongDesc,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires an argument")
			}
			return nil
		},
		Run:        runDestroy,
		SuggestFor: []string{"delete", "remove"},
	}

	cmd.SetHelpTemplate(t.HelpTemplate)
	cmd.SetUsageTemplate(t.UsageTemplate)

	return cmd
}

func runDestroy(cmd *cobra.Command, args []string) {
	ui.ProgressBar("Querying state", 1500, false)
	if !noop {
		ui.DescribeStep("OK, here's the plan. Lyra will destroy the following resources:")
	} else {
		ui.NoopBanner()
		ui.DescribeStep("OK, here's the plan. Lyra would attempt to perform the following actions:")
	}
	fmt.Println() // This is stupid, but it needs some padding
	ui.DiffRemove(`lyra::aws::ec2_instance:
    myapp-vm-1:
      region:              eu-west-1
      availability_zone:   eu-west-1a
      image_id:            ami-b63ae0ce
      instance_type:       t2.micro
      key_name:            auth
      security_groups:
        default
      tags:
        environment: dev_env
`)

	ui.DiffRemove(`lyra::aws::ec2_instance:
    myapp-vm-2:
      region:              eu-west-1
      availability_zone:   eu-west-1a
      image_id:            ami-b63ae0ce
      instance_type:       t2.micro
      key_name:            auth
      security_groups:
        default
      tags:
        environment: dev_env
`)

	if !noop {
		c := ui.AskForConfirmation("Sure you want to do this?")
		if c {
			ui.ProgressBar("Destroying ec2_instance 'myapp-vm-1'", 1500, false)
			ui.ResourceDestroy("ec2_instance 'myapp-vm-1'")
			if chaos {
				ui.ResourceError("Could not destroy ec2_instance 'myapp-vm-2'")
				ui.Failure("Terrible things happened")
			} else {
				ui.ProgressBar("Destroying ec2_instance 'myapp-vm-2'", 1500, false)
				ui.ResourceDestroy("ec2_instance 'myapp-vm-2'")
				m.UnsetDirty()
				ui.Success("Destroy finished without error")
				ui.ResourceSummary("0", "2", "0", "0")
			}
			return
		}
	} else {
		ui.NoopBanner()
	}
}
