package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/lyraproj/lyra/pkg/i18n"
	"github.com/lyraproj/lyra/pkg/logger"
	"github.com/lyraproj/lyra/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultLogEncoding = "console"
)

var (
	assumeYes bool
	chaos     bool
	debug     bool
	jsonlogs  bool
	loglevel  string
	noop      bool
	outline   bool
	verbose   bool
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

	cmd.PersistentFlags().BoolVarP(&assumeYes, "assume-yes", "y", false, "Bypass any y/n prompts, answering yes")
	cmd.PersistentFlags().BoolVarP(&chaos, "chaos", "c", false, "")
	cmd.PersistentFlags().BoolVar(&debug, "debug", false, i18n.T("rootFlagDebug"))
	cmd.PersistentFlags().BoolVar(&jsonlogs, "jsonlogs", false, "Output logs in JSON format")
	cmd.PersistentFlags().StringVar(&loglevel, "loglevel", "", i18n.T("rootFlagLoglevel"))
	cmd.PersistentFlags().BoolVar(&noop, "dry-run", false, "Dry-run, no-op mode")
	cmd.PersistentFlags().BoolVarP(&outline, "outline", "o", false, "")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Show verbose output")
	// Hide internal/dev flags from users
	cmd.PersistentFlags().MarkHidden("outline")
	cmd.PersistentFlags().MarkHidden("chaos")

	// Some (all? most?) flags should be accessable in config file
	// TODO: Would be nice to put CLI stuff in something other than the top level
	viper.BindPFlag("jsonlogs", cmd.PersistentFlags().Lookup("jsonlogs"))
	viper.BindPFlag("loglevel", cmd.PersistentFlags().Lookup("loglevel"))
	viper.BindPFlag("assume-yes", cmd.PersistentFlags().Lookup("assume-yes"))

	// real
	cmd.AddCommand(NewVersionCmd())
	// cmd.AddCommand(NewApplyCmd())
	cmd.AddCommand(NewControllerCmd())
	cmd.AddCommand(NewValidateCmd())
	cmd.AddCommand(EmbeddedPluginCmd())

	// mock
	cmd.AddCommand(NewApplyMockCmd())
	cmd.AddCommand(NewApplyEnvCmd())
	cmd.AddCommand(NewDestroyCmd())
	cmd.AddCommand(NewExperimentalCmd())
	cmd.AddCommand(NewInitCmd())
	cmd.AddCommand(NewShowCmd())
	cmd.AddCommand(NewCreateCmd())
	cmd.AddCommand(NewRegisterCmd())
	cmd.AddCommand(NewLintCmd())

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

	// TODO: Actually handle this in a meaningful way
	user, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = viper.ReadInConfig() // Find and read the config file

	// FIXME: This feels abhorrent, but how else to ensure Viper
	// config affects the logger? Not sure how to set jsonlogs
	// after initialisation.
	spec := logger.Spec{
		Name:   "lyra",
		Level:  viper.GetString("loglevel"),
		JSON:   viper.GetBool("jsonlogs"),
		Output: os.Stderr,
	}

	logger.Initialise(spec)

	log := logger.Get()

	if err != nil {
		// Log the error but keep going
		log.Debug(err.Error())

		// If no config file, make one. First, touch an empty file.
		// WriteConfig() can't do this itself because ¯\_(ツ)_/¯
		cfgfile := filepath.Join(user.HomeDir, ".lyra/config.yaml")  // FIXME: Windows?
		_, err = os.OpenFile(cfgfile, os.O_RDONLY|os.O_CREATE, 0666) // FIXME: Windows?
		if err != nil {
			log.Error(err.Error())
		} else {
			log.Debug("touched new empty config file", cfgfile)
		}

		// Write values, bound to flags via init() above, to new file.
		err = viper.WriteConfig()
		if err != nil {
			log.Error(err.Error())
		} else {
			log.Debug("wrote values to config file", cfgfile)
		}
	} else {
		log.Debug("success in reading config from", viper.ConfigFileUsed())
	}
	// Log all settings for debug posterity
	c := viper.AllSettings()
	f, err := json.Marshal(c)
	if err != nil {
		log.Error("unable to marshal config to JSON: %v", err)
	}
	log.Debug("current config values", string(f))
}
