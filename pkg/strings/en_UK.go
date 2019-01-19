package strings

// Strings for UI
const (
	exampleManifest  = "manifest.pp"
	examplePluginDir = "~/.lyra/plugins"
	exampleTemplate  = "aws-ec2"
	exampleDeploy    = "production"
	exampleTask      = "migrate_db"
	exampleResource  = "aws::lyra::ec2_instance[instance-foo]"
	exampleCloud     = "aws-prod"
	examplePadding   = "  " //Two spaces padding to match indentation of other sections. ʕノ•ᴥ•ʔノ ︵ ┻━┻

	sourceLabel = "project"    //Or: "module", "package"
	deployLabel = "deployment" //Or: "stack", "application"

	RootCmdName      = "lyra"
	RootCmdShortDesc = "Create and configure cloud infrastructure and resources"
	RootCmdLongDesc  = "Lyra creates and configures cloud infrastructure and resources"

	ApplyCmdName      = "apply"
	ApplyCmdUsage     = ApplyCmdName + " <manifest>"
	ApplyCmdExample   = examplePadding + RootCmdName + " " + ApplyCmdName + " " + exampleManifest
	ApplyCmdShortDesc = "Create or update infrastructure resources"
	ApplyCmdLongDesc  = "Create or update infrastructure resources"

	DestroyCmdName    = "destroy"
	DestroyCmdUsage   = DestroyCmdName + " <" + deployLabel + "|resource>"
	DestroyCmdExample = examplePadding + RootCmdName + " " + DestroyCmdName + " " + exampleDeploy + "\n" +
		examplePadding + RootCmdName + " " + DestroyCmdName + " " + exampleResource
	DestroyCmdShortDesc = "Destroy infrastructure resources"
	DestroyCmdLongDesc  = "Destroy infrastructure resources"

	CreateCmdName      = "create"
	CreateCmdUsage     = CreateCmdName + " <template name>"
	CreateCmdExample   = examplePadding + RootCmdName + " " + CreateCmdName + " " + exampleTemplate
	CreateCmdShortDesc = "Create a new Lyra workflow package"
	CreateCmdLongDesc  = "Create a new Lyra workflow package"

	PlanCmdName      = "plan"
	PlanCmdUsage     = PlanCmdName + " <" + sourceLabel + ">"
	PlanCmdExample   = examplePadding + RootCmdName + " " + PlanCmdName + " " + exampleDeploy
	PlanCmdShortDesc = "Parse and evaluate, but do not apply, a manifest and data"
	PlanCmdLongDesc  = "Parse and evaluate, but do not apply, a manifest and data"

	PluginCmdName      = "plugin"
	PluginCmdUsage     = PluginCmdName + " /path/to/plugindir"
	PluginCmdExample   = examplePadding + RootCmdName + " " + PluginCmdName + " " + examplePluginDir
	PluginCmdShortDesc = "List plugins found in supplied discovery directory"

	RunCmdName      = "run"
	RunCmdUsage     = RunCmdName + " <task name> <target resource>"
	RunCmdExample   = examplePadding + RootCmdName + " " + RunCmdName + " " + exampleTask + " " + exampleResource
	RunCmdShortDesc = "Execute an imperative task against cloud resources"
	RunCmdLongDesc  = "Execute an imperative task against cloud resources"

	ShowCmdName  = "show"
	ShowCmdUsage = ShowCmdName + " <" + deployLabel + "|cloud|type|resource>"
	// FIXME: This ↓ is stupid. Do something better.
	ShowCmdExample = examplePadding + RootCmdName + " " + ShowCmdName + " " + exampleCloud + "\n" +
		examplePadding + RootCmdName + " " + ShowCmdName + " lyra::aws::ec2_instance\n" +
		examplePadding + RootCmdName + " " + ShowCmdName + " " + exampleDeploy + "\n" +
		examplePadding + RootCmdName + " " + ShowCmdName + " lyra::aws::ec2_instance['myapp-vm-1']\n"
	ShowCmdShortDesc = "Enumerate extant cloud resources"
	ShowCmdLongDesc  = "Enumerate extant cloud resources"

	ValidateCmdName      = "validate"
	ValidateCmdUsage     = ValidateCmdName + " <manifest>"
	ValidateCmdExample   = examplePadding + RootCmdName + " " + ValidateCmdName + " " + exampleManifest
	ValidateCmdShortDesc = "Validates syntax"
	ValidateCmdLongDesc  = "Validates syntax"

	VersionCmdName      = "version"
	VersionCmdUsage     = VersionCmdName
	VersionCmdShortDesc = "Version of the Lyra client"
	VersionCmdExample   = examplePadding + RootCmdName + " " + VersionCmdName
)

// HelpTemplate is helpful
// Inspired by https://github.com/kubernetes/kompose/blob/master/cmd/convert.go
// Remember ALL the whitespace is significant!
var HelpTemplate = `Description:
  {{rpad .Long 10}}

Usage:{{if .Runnable}}
{{if .HasAvailableFlags}}  {{appendIfNotPresent .UseLine "[flags]"}}{{else}}{{.UseLine}}{{end}}{{end}}
{{if gt .Aliases 0}}

Aliases:
  {{.NameAndAliases}}
{{end}}{{if .HasExample}}

Examples:
{{ .Example }}{{end}}{{ if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimRightSpace}}{{end}}{{ if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimRightSpace}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsHelpCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableSubCommands }}{{end}}
`

// UsageTemplate is similar to HelpTemplate, but sticks to brief usage and examples.
// Remember ALL the whitespace is significant!
var UsageTemplate = `
Usage:{{if .Runnable}}
{{if .HasAvailableFlags}}  {{appendIfNotPresent .UseLine "[flags]"}}{{else}}{{.UseLine}}{{end}}{{end}}{{if gt .Aliases 0}}

Aliases:
{{.NameAndAliases}}
{{end}}{{if .HasExample}}

Examples:
{{ .Example }}{{end}}{{ if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableLocalFlags}}

See '{{.CommandPath}} --help' for help and examples.{{end}}`
