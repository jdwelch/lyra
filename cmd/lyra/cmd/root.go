package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lyraproj/lyra/cmd/lyra/ui"
	"github.com/lyraproj/lyra/pkg/i18n"
	"github.com/lyraproj/lyra/pkg/logger"
	"github.com/lyraproj/lyra/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const (
	defaultLogEncoding = "console"
)

var (
	assumeYes         bool
	chaos             bool
	connectiontimeout int
	debug             bool
	host              string
	jsonlogs          bool
	kubeconfig        string
	kubecontext       string
	kubenamespace     string
	loglevel          string
	noop              bool
	outline           bool
	verbose           bool
)

// NewRootCmd returns the root command
func NewRootCmd() *cobra.Command {

	// Set up gettext
	i18n.Configure("locales", "en_US", "default")

	cmd := &cobra.Command{
		Use:              i18n.T("rootCmdUse"),
		Short:            i18n.T("rootCmdShort"),
		Long:             i18n.T("rootCmdLong"),
		Run:              runHelp,
		PersistentPreRun: initialiseTool,
		Version:          fmt.Sprintf("%v", version.Get()),
	}

	// Flags for controlling output
	cmd.PersistentFlags().BoolVar(&debug, "debug", false, i18n.T("rootFlagDebug"))
	cmd.PersistentFlags().BoolVar(&jsonlogs, "jsonlogs", false, "Output logs in JSON format")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Show verbose output")
	cmd.PersistentFlags().StringVar(&loglevel, "loglevel", "", i18n.T("rootFlagLoglevel"))

	// Flags for controlling behaviour
	cmd.PersistentFlags().BoolVar(&noop, "dry-run", false, "Dry-run  mode")
	cmd.PersistentFlags().BoolVarP(&assumeYes, "assume-yes", "y", false, "Bypass any y/n prompts, answering yes")

	// Flags for k8s
	if home := homedir.HomeDir(); home != "" { // Not sure why I have to do this, but if I don't kube client complains later.
		cmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", filepath.Join(home, ".kube", "config"), "absolute path to the kubeconfig file to use")
	} else {
		cmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file to use")
	}
	cmd.PersistentFlags().IntVarP(&connectiontimeout, "cygnus-connection-timeout", "t", 300, "the duration (in seconds) Lyra will wait to establish a connection to Cygnus")
	cmd.PersistentFlags().StringVar(&host, "cygnus-host", getKubeHost(kubeconfig), "Address of Cygnus (probably a Kubernetes master)")
	cmd.PersistentFlags().StringVarP(&kubenamespace, "kube-namespace", "n", "kube-system", "Kubernetes namespace for Cygnus")
	cmd.PersistentFlags().StringVar(&kubecontext, "kube-context", "", "Name of the kubeconfig context to use")

	cmd.SetHelpTemplate(ui.HelpTemplate)
	cmd.SetUsageTemplate(ui.UsageTemplate)

	// Flags for design purposes
	cmd.PersistentFlags().BoolVarP(&outline, "outline", "o", false, "")
	cmd.PersistentFlags().MarkHidden("outline")
	cmd.PersistentFlags().BoolVarP(&chaos, "chaos", "c", false, "")
	cmd.PersistentFlags().MarkHidden("chaos")

	// Most flags should be accessable in config file
	viper.BindPFlag("dry-run", cmd.PersistentFlags().Lookup("dry-run"))
	viper.BindPFlag("jsonlogs", cmd.PersistentFlags().Lookup("jsonlogs"))
	viper.BindPFlag("loglevel", cmd.PersistentFlags().Lookup("loglevel"))
	viper.BindPFlag("assume-yes", cmd.PersistentFlags().Lookup("assume-yes"))
	viper.BindPFlag("cygnus-host", cmd.PersistentFlags().Lookup("cygnus-host"))
	viper.BindPFlag("cygnus-connection-timeout", cmd.PersistentFlags().Lookup("cygnus-connection-timeout"))
	viper.BindPFlag("kubeconfig", cmd.PersistentFlags().Lookup("kubeconfig"))
	viper.BindPFlag("kube-context", cmd.PersistentFlags().Lookup("kube-context"))
	viper.BindPFlag("kube-namespace", cmd.PersistentFlags().Lookup("kube-namespace"))

	// Commands
	cmd.AddCommand(NewApplyEnvCmd())
	cmd.AddCommand(NewApplyMockCmd())
	cmd.AddCommand(NewControllerCmd())
	cmd.AddCommand(NewCreateCmd())
	cmd.AddCommand(NewDescribeCmd())
	cmd.AddCommand(NewDestroyCmd())
	cmd.AddCommand(NewExperimentalCmd())
	cmd.AddCommand(NewInitCmd())
	cmd.AddCommand(NewLintCmd())
	cmd.AddCommand(NewRegisterCmd())
	cmd.AddCommand(NewShowCmd())
	cmd.AddCommand(NewStatusCmd())
	cmd.AddCommand(NewUninstallCmd())
	cmd.AddCommand(NewUnregisterCmd())
	cmd.AddCommand(NewUpgradeCmd())
	cmd.AddCommand(NewValidateCmd())
	cmd.AddCommand(NewVersionCmd())

	return cmd
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func initialiseTool(cmd *cobra.Command, args []string) {
	if debug {
		loglevel = "debug"
	}

	// Set up user configuration & logger
	initializeConfig()
}

// Set up persistent config via Viper
// TODO: Let user specify config file location
func initializeConfig() {

	viper.SetConfigName("config")      // name of config file (without extension)
	viper.AddConfigPath("/etc/lyra/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.lyra") // call multiple times to add many search paths
	viper.AddConfigPath(".")           // optionally look for config in the working directory
	viper.SetConfigType("yaml")

	spec := logger.Spec{
		Name:   "lyra",
		Level:  viper.GetString("loglevel"),
		JSON:   viper.GetBool("jsonlogs"),
		Output: os.Stderr,
	}
	logger.Initialise(spec)
	log := logger.Get()

	// Read config from disk
	// FIXME: WTFF
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Debug("success in reading config from", viper.ConfigFileUsed())
	// } else {
	// 	log.Error(err.Error())
	// }

	// Log current settings for debug posterity
	c := viper.AllSettings()
	f, err := json.Marshal(c)
	if err != nil {
		log.Error("unable to marshal config to JSON: %v", err)
	}
	log.Debug("current config values", string(f))
}

func getKubeHost(kubeconfig string) string {
	// https://github.com/kubernetes/client-go/blob/master/examples/create-update-delete-deployment/main.go
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	if config.Host != "" {
		return config.Host
	}
	return ""
}
