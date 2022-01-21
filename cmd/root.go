package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"

	mainApp "github.com/vasylNaumenko/ej/internal/app"
	"github.com/vasylNaumenko/ej/internal/domain/config"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ej",
	Short: "An easy jira task creator with notifications to the Discord",
	Long:  `An easy jira task creator with notifications to the Discord`,
}

const (
	appKey = "app"
	cfgKey = "cfg"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(
		initConfig,
		initApp,
	)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ej.config.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".ej.config" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ej.config")
		viper.SetConfigType("yaml")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// Parse config file.
	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if errs := cfg.Validate(); len(errs) != 0 {
		fmt.Println(fmt.Errorf("validation error - %s", strings.Join(errs, ",")))
		os.Exit(1)
	}

	viper.Set(cfgKey, cfg)
}

// initApp is intends to initialize tha main application layer
func initApp() {
	app, err := mainApp.NewApp(viper.Get(cfgKey).(config.Config))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.Set(appKey, app)
}

// getApp returns the main application or fails
func getApp() *mainApp.App {
	app, ok := viper.Get(appKey).(*mainApp.App)
	if !ok {
		fmt.Println("can`t get the App")
		os.Exit(1)
	}
	return app
}

// exitIfError echoes an error and halts the application.
func exitIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
