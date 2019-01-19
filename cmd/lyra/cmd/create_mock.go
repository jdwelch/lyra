package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/lyraproj/lyra/cmd/lyra/ui"
	"github.com/lyraproj/lyra/pkg/logger"
	m "github.com/lyraproj/lyra/pkg/mock"
	t "github.com/lyraproj/lyra/pkg/strings"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// LyraPlugin is metadata about the project
type LyraPlugin struct {
	Name        string `survey:"name"`
	Template    string `survey:"platform"`
	Language    string `survey:"language"`
	LanguageExt string
}

// NewCreateCmd returns the create subcommand
func NewCreateCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:        t.CreateCmdName,
		Example:    t.CreateCmdExample,
		Short:      t.CreateCmdShortDesc,
		Long:       t.CreateCmdLongDesc,
		Run:        runCreate,
		SuggestFor: []string{"new"},
	}

	cmd.SetHelpTemplate(t.HelpTemplate)
	cmd.SetUsageTemplate(t.UsageTemplate)

	return cmd
}

func runCreate(cmd *cobra.Command, args []string) {

	name := "my-project"

	if len(args) > 0 {
		name = args[0]
	}

	log := logger.Get()
	var qs = []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "What would you like to call this Lyra plugin?",
				Default: name,
			},
			// FIXME: Also validate for proper format
			Validate: func(val interface{}) error {
				if m.Exists(val.(string)) {
					return errors.New("directory " + val.(string) + " exists. Try another name, please.")
				}
				return nil
			},
			Transform: survey.ToLower,
		},
		{
			Name: "template",
			Prompt: &survey.Select{
				Message: "Choose a template:",
				Options: []string{
					// TODO: Read these from a repo somewhere else
					"aws-iaas",
					"faas (cloud-agnostic)",
					"google kubernetes engine",
				},
				Default: "aws",
			},
		},
		{
			Name: "language",
			Prompt: &survey.Select{
				Message: "Choose a language:",
				// FIXME: Translate these to file extensions
				Options: []string{"puppet", "yaml", "typescript"},
				Default: "puppet",
			},
		},
	}

	answers := LyraPlugin{
		`survey: "name"`,
		`survey: "template"`,
		`survey: "language"`,
		".ext",
	}

	ui.DescribeStep("Lyra can help you get started with a helpful scaffold.")
	fmt.Println("Just answer these questions:")

	err := survey.Ask(qs, &answers)
	if err != nil {
		log.Error("No answer", "error", err)
		os.Exit(0)
	}

	createManifestScaffold(answers)

}

func createManifestScaffold(stack LyraPlugin) {

	// FIXME: This seems stupid.
	if stack.Language == "puppet" {
		stack.LanguageExt = "pp"
	}

	if stack.Language == "yaml" {
		stack.LanguageExt = "yaml"
	}

	if stack.Language == "typescript" {
		stack.LanguageExt = "ts"
	}

	modulename := strings.ToLower(stack.Name)
	moduledirectory := strings.ToLower(stack.Name)

	manifestdir := moduledirectory + ""
	manifestfile := manifestdir + "/" + modulename + "-manifest." + stack.LanguageExt

	datadir := moduledirectory + ""
	datafile := datadir + "/" + modulename + "-vars.yaml"

	taskdir := moduledirectory + "/tasks"

	mkScaffoldDir(moduledirectory)
	mkScaffoldDir(datadir)
	mkScaffoldDir(manifestdir)
	mkScaffoldDir(taskdir)

	// This is stupid
	// FIXME: Stop doing stupid
	if stack.Language == "puppet" {
		generateFileFromTemplate(stack, manifestfile, puppetTemplate)
	}
	if stack.Language == "yaml" {
		generateFileFromTemplate(stack, manifestfile, yamlTemplate)
	}
	if stack.Language == "typescript" {
		generateFileFromTemplate(stack, manifestfile, typescriptTemplate)
	}

	generateFileFromTemplate(stack, datafile, yamlDataTemplate)

	ui.Success("Created a new Lyra project scaffold with this structure:")

	showModuleStructure(stack)

	fmt.Println("\nYour project is ready to use. Run 'lyra apply " + manifestfile + " --data " + datafile + "'\n")

}

func mkScaffoldDir(path string) {
	log := logger.Get()
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Error(path)
	}
	log.Debug(path)
}

func generateFileFromTemplate(stack LyraPlugin, filename string, filetemplate string) {
	log := logger.Get()
	buf := new(bytes.Buffer)

	file, err := os.Create(filename)
	if err != nil {
		log.Error(err.Error())
	}
	log.Debug("create" + filename)

	defer file.Close()

	tmpl, err := template.New(stack.Name).Parse(filetemplate)
	if err != nil {
		log.Error(err.Error())
	}

	buf.Reset()
	err = tmpl.Execute(buf, stack)
	if err != nil {
		log.Error(err.Error())
	}

	fmt.Fprintf(file, buf.String())
}

func showModuleStructure(stack LyraPlugin) {
	log := logger.Get()
	buf := new(bytes.Buffer)

	tmpl, err := template.New(stack.Name).Parse(directoryTree)
	if err != nil {
		log.Error(err.Error())
	}

	buf.Reset()
	err = tmpl.Execute(buf, stack)
	if err != nil {
		log.Error(err.Error())
	}

	fmt.Printf(buf.String())
}

// FIXME: Move this into proper files and stuff
// Or not. You get the idea either way.
const (
	puppetTemplate = `# This is an auto-generated scaffold of a Lyra plugin
# for Amazon Web Services. For detailed documentation,
# see https://github.com/lyraproj/lyra/docs/getting-started.md

lyra::aws::instance { '{{.Name}}-instance':
  ensure             => 'present',
  region             => lookup('aws_region'),
  image_id           => lookup('image_id'),
  instance_type      => lookup't2.micro',
  key_name           => lookup(env.'my_aws_key'),
  tag_specifications => [ Lyra::Aws::Tagspecification ({
    'resource_type' => 'instance',
    'tags'          => {
      'created_by' => 'admin@example.com',
      'department' => 'engineering',
      'project'    => 'incubator',
    }
  }) ],
} -> notice("Created EC2 Instance, ID: ${resource(Lyra::Aws::Instance['{{.Name}}-instance']).instance_id}, PrivateIP: ${resource(Lyra::Aws::Instance['{{.Name}}-instance']).private_ip_address}, PublicIP: ${resource(Lyra::Aws::Instance['{{.Name}}-instance']).public_ip_address}\n")

$install_nginx = @("INSTALL_NGINX"/L)
sudo apt -y update
sudo apt -y install nginx
sudo rm -rf /var/www/html/index.html
echo "<html><head><title>Installed by Lyra</title></head><body><h1>INSTALLED BY LYRA</h1></body></html>" | sudo tee /var/www/html/index.html
| INSTALL_NGINX

lyra::ssh::exec { '{{.Name}}-exec':
  ensure  => present,
  host    => resource(Lyra::Aws::Instance['{{.Name}}-instance']).public_ip_address,
  user    => 'ubuntu',
  command => $install_nginx,
  port    => 22,
} ->
notice("Installed NGinx: http://${resource(Lyra::Aws::Instance['{{.Name}}-instance']).public_ip_address}\n")
`

	yamlTemplate = `sequential:
  foo:
    bar:
`
	typescriptTemplate = `var tier1 = function() {
  {{.Name}}
  resources({
    'lyra::aws::importkeypair': {
      'myapp-keypair': {
        ensure: 'present',
        region: region,
        public_key_material: public_key
      }
    },
  });
`

	yamlDataTemplate = `# This is an auto-genereated scaffold
# of a Lyra data file. For detailed documentation,
# see https://github.com/lyraproj/lyra/docs/getting-started.md
---
aws_region: 'us-west-2'
image_id: 'ami-b63ae0ce'
`

	directoryTree = `{{.Name}}
├── {{.Name}}-manifest.{{.LanguageExt}}
├── {{.Name}}-vars.yaml
└── tasks
`
)
