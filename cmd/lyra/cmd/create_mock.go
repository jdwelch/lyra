package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"os/user"
	"strings"

	"github.com/lyraproj/lyra/cmd/lyra/ui"
	"github.com/lyraproj/lyra/pkg/logger"
	m "github.com/lyraproj/lyra/pkg/mock"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// LyraPlugin is metadata about the project
type LyraPlugin struct {
	Name        string `survey:"name"`
	Author      string
	Version     string
	Language    string `survey:"language"`
	LanguageExt string
}

// NewCreateCmd returns the create subcommand
func NewCreateCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:        "create",
		Example:    "create",
		Short:      "Create a new Lyra workflow package",
		Long:       "Create a new Lyra workflow package",
		Run:        runCreate,
		SuggestFor: []string{"new"},
	}

	cmd.SetHelpTemplate(ui.HelpTemplate)
	cmd.SetUsageTemplate(ui.UsageTemplate)

	return cmd
}

func runCreate(cmd *cobra.Command, args []string) {

	name := "my-project"
	author := whoAmI()
	version := "0.1.0"
	language := "puppet"

	if len(args) > 0 {
		name = args[0]
	}

	log := logger.Get()
	var qs = []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "What would you like to call this Lyra package?",
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
			Name: "language",
			Prompt: &survey.Select{
				Message: "Choose a language:",
				Options: []string{"puppet", "yaml", "typescript"},
				Default: language,
			},
		},
	}

	answers := LyraPlugin{
		`survey: "name"`,
		author,
		version,
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

func createManifestScaffold(wf LyraPlugin) {

	// Generate package directory
	pkgname := strings.ToLower(wf.Name)
	pkgdirectory := strings.ToLower(wf.Name)
	mkScaffoldDir(pkgdirectory)

	// Generate metadata
	metadatafile := pkgdirectory + "/" + "metadata.yaml"
	generateFileFromTemplate(wf, metadatafile, metadataTemplate)

	// Generate workflow template
	wfdir := pkgdirectory + "/workflows"
	mkScaffoldDir(wfdir)

	// FIXME: This manual mapping seems stupid.
	if wf.Language == "puppet" {
		wf.LanguageExt = "pp"
	}

	if wf.Language == "yaml" {
		wf.LanguageExt = "yaml"
	}

	if wf.Language == "typescript" {
		wf.LanguageExt = "ts"
	}

	wffile := wfdir + "/" + pkgname + "-default." + wf.LanguageExt

	if wf.Language == "puppet" {
		generateFileFromTemplate(wf, wffile, puppetWfTemplate)
	}
	if wf.Language == "yaml" {
		generateFileFromTemplate(wf, wffile, yamlWfTemplate)
	}
	if wf.Language == "typescript" {
		generateFileFromTemplate(wf, wffile, typescriptWfTemplate)
	}

	// Generate sample data
	datafile := pkgdirectory + "/" + "values.yaml"
	generateFileFromTemplate(wf, datafile, dataTemplate)

	// Generate sample k8s object spec
	objectspec := pkgdirectory + "/" + "deploy.yaml"
	generateFileFromTemplate(wf, objectspec, specTemplate)

	// Generate readme
	readme := pkgdirectory + "/" + "README.md"
	generateFileFromTemplate(wf, readme, readmeTemplate)

	ui.Success("Created a new Lyra project scaffold with this structure:")
	showPkgStructure(wf)

	fmt.Println("\nYour workflow is ready to use. Run 'lyra apply " + wffile + " --data " + datafile + "'\n")
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

func showPkgStructure(wf LyraPlugin) {
	log := logger.Get()
	buf := new(bytes.Buffer)

	tmpl, err := template.New(wf.Name).Parse(directoryTree)
	if err != nil {
		log.Error(err.Error())
	}

	buf.Reset()
	err = tmpl.Execute(buf, wf)
	if err != nil {
		log.Error(err.Error())
	}

	fmt.Printf(buf.String())
}

func whoAmI() string {
	log := logger.Get()
	user, err := user.Current()
	if err != nil {
		log.Error(err.Error())
	}
	return user.Username
}

const (

	// FIXME: Make this match what actually happens instead of faking it so badly
	directoryTree = `{{.Name}}
├── README.md
├── metadata.yaml
├── deploy.yaml
├── values.yaml
└── workflows
    └──{{.Name}}-default.{{.LanguageExt}}
`

	puppetWfTemplate = `# This is an auto-generated scaffold
# of a Lyra workflow. For detailed documentation,
# see https://github.com/lyraproj/lyra/docs/getting-started.md

workflow sample2 {
  typespace => 'example',
  input => (
    String $foo = lookup('foo', undef, undef, "foo"),
    String $bar = lookup('bar', undef, undef, "bar"),
    String $baz = lookup('baz', undef, undef, "baz")
  ),
  output => (
    String $foo,
    String $bar,
    String $baz
  )
} {
  resource person {
    output => ($name)
  }{
    age => 28,
    name => 'Bob',
    human => false,
  }
}
`

	yamlWfTemplate = `# This is an auto-generated scaffold
# of a Lyra workflow. For detailed documentation,
# see https://github.com/lyraproj/lyra/docs/getting-started.md
# ---
# TODO: WORKFLOW (in YAML) HERE!
`

	typescriptWfTemplate = `// This is an auto-generated scaffold
// of a Lyra workflow. For detailed documentation,
// see https://github.com/lyraproj/lyra/docs/getting-started.md
// TODO: WORKFLOW (in TypeScript) HERE!
`

	dataTemplate = `# This is an auto-genereated scaffold
# of a Lyra data file. For detailed documentation,
# see https://github.com/lyraproj/lyra/docs/getting-started.md
---
aws_region: 'us-west-2'
image_id: 'ami-b63ae0ce'
`

	metadataTemplate = `# This is an auto-genereated scaffold
# of a Lyra package metadata file. For detailed documentation,
# see https://github.com/lyraproj/lyra/docs/getting-started.md
---
apiVersion: v1alpha1
name: {{.Name}}
author: {{.Author}}
description: "A lovely Lyra workflow."
version: {{.Version}}
license: "Apache 2"
url: "gh.com/foo/bar"
`
	readmeTemplate = `# README for {{.Name}} goes here!`

	specTemplate = `# This is an auto-genereated scaffold
# of a Lyra package deployment file. For detailed documentation,
# see https://github.com/lyraproj/lyra/docs/getting-started.md
---
apiVersion: lyraproj.io/v1alpha1
kind: Workflow
metadata:
  app.kubernetes.io/name: {{.Name}}-wf
  labels:
    lyraproj.io/workflowVersion: {{.Version}}
spec:
  lyraproj.io/workflowName: "{{.Name}}-wf"
	lyraproj.io/data:
	  aws_region: 'us-west-2'
    image_id: 'ami-b63ae0ce'
    refreshTime: 60
`
)
