package cmd

import (
	"fmt"
	"os"
	"strings"

	ui "github.com/lyraproj/lyra/cmd/lyra/ui"
	"github.com/lyraproj/lyra/pkg/i18n"
	"github.com/lyraproj/lyra/pkg/logger"
	m "github.com/lyraproj/lyra/pkg/mock"
	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	lang      string
	extData   string
	hieraData string
)

// NewApplyMockCmd returns the apply (mock) subcommand
// Yeah, I should just base this on the actual apply subcommand
func NewApplyMockCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     i18n.T("applyCmdUse"),
		Short:   i18n.T("applyCmdShort"),
		Long:    i18n.T("applyCmdLong"),
		Example: i18n.T("applyCmdExample"),
		Run:     runApplyMock,
		Args:    cobra.ExactArgs(1),
	}
	cmd.Flags().StringVarP(&homeDir, "root", "r", "", "path to root directory")
	cmd.Flags().StringVarP(&hieraData, "data", "d", "./data.yaml", "path to hiera data")
	return cmd
}

func runApplyMock(cmd *cobra.Command, args []string) {
	log := logger.Get()
	log.Debug("applying manifest")
	// Get current values from config file
	// TODO: Why does these have to live here?
	// I clearly don't understand the command lifecycle
	assumeYes = viper.GetBool("assume-yes")

	path := args[0]

	if outline {
		runApplyMockOutline(cmd, args)
		os.Exit(0)
	}

	log.Debug("reading manifest", "path", path)
	ui.DescribeStepWithField("Applying manifest:", path)

	if viper.GetBool("apply-dirty") == false {
		mockApplyFirstTime()
	} else {
		mockApplyChange()
	}

	log.Debug("apply finished")
}

func mockApplyFirstTime() {
	if !noop {
		fmt.Println("OK, here's the plan. Lyra will perform the following actions:")
	} else {
		ui.NoopBanner()
		fmt.Println("OK, here's the plan. Lyra would attempt to perform the following actions:")
	}
	fmt.Println() // This is stupid, but it needs some padding
	ui.DiffAdd(`lyra::aws::ec2_instance:
    'myapp-vm-1':
      region: eu-west-1
      image_id: ami-b63ae0ce
      instance_type: t2.micro
      key_name: id_rsa
      security_groups:
        default
      tags:
        created_by: admin@example.com
        department: engineering
        project: incubator
`)

	ui.DiffAdd(`lyra::ssh::exec:
    'myapp-exec':
      ensure: present
      host: resource(Lyra::Aws::Instance['{{.Name}}-instance']).public_ip_address
      user: 'ubuntu'
      command: 'sudo apt -y update\nsudo apt -y install nginx\nsudo rm -rf /var/www/html/index.html\necho <html><head><title>Installed by Lyra</title></head><body><h1>INSTALLED BY LYRA</h1></body></html>" | sudo tee /var/www/html/index.html'
      port: 22
`)

	if !noop {
		c := ui.AskForConfirmation("Sure you want to do this?")
		if c {
			m.SetDirty()
			ui.ProgressBar("Creating ec2_instance myapp-vm-1", 1500, true)
			if !verbose {
				ui.ResourceSet("Created ec2_instance myapp-vm-1")
			} else {
				ui.ResourceSet("\n" + m.VerboseResouce)
			}
			ui.Notice("EC2 Instance ID: i-0c73357d0430ea0fe, PrivateIP: 172.31.29.197, PublicIP: 54.171.56.145")
			ui.ProgressBar("Creating ec2_instance myapp-vm-2", 1500, false)
			if chaos {
				ui.ResourceError("Could not create ec2_instance myapp-vm-2")
				ui.Failure("Terrible things happened")
			} else {
				if !verbose {
					ui.ResourceSet("Created ec2_instance myapp-vm-2")
				} else {
					ui.ResourceSet("\n" + m.VerboseResouce)
				}
				ui.Notice("EC2 Instance ID: i-3a2f357dea0fe0430, PrivateIP: 172.31.29.197, PublicIP: 37.171.23.125")
				ui.Notice("Installed NGinx: http://37.171.23.125")
				ui.Success("Apply finished without error")
				ui.ResourceSummary("2", "0", "0", "0")
			}
			return
		}
	} else {
		ui.NoopBanner()
	}
}

func mockApplyChange() {
	ui.ProgressBar("Reading remote state", 1500, true)
	if !noop {
		fmt.Println("OK, here's the plan. Lyra will perform the following actions:")
	} else {
		ui.NoopBanner()
		fmt.Println("OK, here's the plan. Lyra would attempt to perform the following actions:")
	}
	fmt.Println() // This is stupid, but it needs some padding
	ui.DiffUnchanged(`lyra::aws::ec2_instance:
    'myapp-vm-1':
      availability_zone:   eu-west-1a
      image_id:            ami-b63ae0ce`)
	ui.DiffRemove("      instance_type:     t2.micro")
	ui.DiffAdd("      instance_type:     m3.medium")
	ui.DiffUnchanged(`      key_name:          auth
      security_groups:
        default
      tags:
        environment: dev_env
`)
	ui.DiffUnchanged(`lyra::aws::ec2_instance:
    'myapp-vm-1':
      availability_zone:   eu-west-1a
      image_id:            ami-b63ae0ce`)
	ui.DiffRemove("      instance_type:     t2.micro")
	ui.DiffAdd("      instance_type:     m3.medium")
	ui.DiffUnchanged(`      key_name:          auth
      security_groups:
        default
      tags:
        environment: dev_env`)
	if !noop {
		c := ui.AskForConfirmation("Sure you want to do this?")
		if c {
			ui.ProgressBar("Changing attributes of ec2_instance myapp-vm-1", 1500, false)
			if !verbose {
				ui.ResourceSet("Changed attributes of instance myapp-vm-1")
			} else {
				ui.ResourceSet("\n" + m.VerboseResouce)
			}
			ui.Notice("EC2 Instance ID: i-0c73357d0430ea0fe, PrivateIP: 172.31.29.197, PublicIP: 54.171.56.145")
			if chaos {
				ui.ResourceError("Could not edit ec2_instance 'myapp-vm-2'")
				ui.DiffConflict("ec2_instance 'myapp-vm-2'")
				ui.DiffConflict("  key: value1 <- yours")
				ui.DiffConflict("ec2_instance 'myapp-vm-2'")
				ui.DiffConflict("  key: value2 <- theirs")
			} else {
				ui.ProgressBar("Changing attributes of e2_instance myapp-vm-2", 1500, false)
				if !verbose {
					ui.ResourceSet("Changed attributes of instance myapp-vm-2")
				} else {
					ui.ResourceSet("\n" + m.VerboseResouce)
				}
				ui.Notice("EC2 Instance ID: i-3a2f357dea0fe0430, PrivateIP: 172.31.29.197, PublicIP: 37.171.23.125")
				ui.Notice("Installed NGinx: http://37.171.23.125")
				ui.Success("Apply finished without error")
				ui.ResourceSummary("0", "0", "2", "0")
			}
		}
	} else {
		ui.NoopBanner()
		return
	}
}

func runApplyMockOutline(cmd *cobra.Command, args []string) {
	path := args[0]

	if noop {
		ui.NoopBanner()
	}
	ui.DescribeStep("[Reading manifest from '" + path + "']")
	m.Outline(m.Process, "[Evaluating manifest]")
	m.Outline(m.Process, " ↳ Incorporate $ENV data")
	if len(getEnv()) > 0 {
		for _, element := range getEnv() {
			variable := strings.Split(element, "=")
			fmt.Println(ansi.Blue+"   "+variable[0], "=", variable[1]+ansi.Reset)
		}
	}
	if extData != "" {
		m.Outline(m.Process, " ↳ Incorporate external data:")
		fmt.Println(ansi.Blue + "   " + extData + ansi.Reset)
	} else {
		m.Outline(m.Process, " ↳ Not using external data")
	}
	m.Outline(m.Process, "[Read remote state…]")
	if !m.Exists(".lyra-apply-dirty") {
		m.Outline(m.Process, " ↳ No existing deployment found")
		ui.DescribeStep("OK, here's the plan:")
		ui.DiffAdd("type[title1]")
		ui.DiffAdd("type[title2]")
		if !assumeYes && !noop {
			m.Outline(m.Question, "Sure you want to do this?")
		}
		if !noop {
			m.Outline(m.Process, "[Attempting to apply resources]")
			ui.ResourceSet("type[title1]")
			ui.Notice("from type[title1]")
			if chaos {
				ui.ResourceError("type[title2]")
				ui.Failure("everything exploded, sorry")
				os.Exit(1)
			} else {
				ui.ResourceSet("type[title2]")
				ui.Notice("from type[title2]")
			}
			ui.DescribeStep("Apply finished")
			m.SetDirty()
		} else {
			ui.DescribeStep("Apply finished")
			ui.NoopBanner()
		}
	} else {
		m.Outline(m.Process, "  ↳ Found existing deployment of these resources")
		ui.DescribeStep("OK, here's what's different:")
		ui.DiffUnchanged("type[title1]")
		ui.DiffAdd("  key: value 2")
		ui.DiffRemove("  key: value 1")
		ui.DiffUnchanged("type[title2]")
		if !assumeYes && !noop {
			m.Outline(m.Question, "Sure you want to do this?")
		}
		if chaos {
			ui.ResourceError("Sorry, there's a problem")
			ui.DiffConflict("type[title2]")
			ui.DiffConflict("  key: value1 <- yours")
			ui.DiffConflict("type[title2.5]")
			ui.DiffConflict("  key: value2 <- theirs")
			m.OutlineSurvey("Which version do you want to keep?", []string{"mine", "theirs", "neither, please bail"})
			os.Exit(1)
		} else {
			ui.ResourceSet("type[title1]")
			ui.Notice("from type[title1]")
			ui.ResourceSet("type[title2]")
			ui.Notice("from type[title2]")
		}
		ui.DescribeStep("Apply finished")
		if noop {
			ui.NoopBanner()
		}
	}
}

func getEnv() []string {
	var vars = []string{}
	for _, elem := range os.Environ() {
		if strings.HasPrefix(elem, "LYRA_") {
			vars = append(vars, elem)
		}
	}
	return vars
}
