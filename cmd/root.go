package cmd

import (
	"fmt"
	"os"

	"github.com/masamerc/sevp/internal"
	"github.com/masamerc/sevp/internal/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "sevp",
	Version: "0.0.4",
	Short:   "sevp: pick and switch environement variables.",
	Long:    `sevp: pick and switch environement variables.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, err := internal.GetConfigFile()
		internal.FailOnError(
			"Error getting config file",
			err,
		)

		contents, err := internal.ReadContents(configPath)
		internal.FailOnError(
			"Error reading config file",
			err,
		)

		profiles := internal.GetProfiles(contents)
		app := app.NewApp(profiles)

		err = app.Run()
		internal.FailOnError(
			"Error running app",
			err,
		)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// TODO: config support

	// configFile = ".cobra-cli-samples.yml"
	// viper.SetConfigType("yaml")
	// viper.SetConfigFile(configFile)
	//
	// viper.AutomaticEnv()
	// viper.SetEnvPrefix("COBRACLISAMPLES")
	// helper.HandleError(viper.BindEnv("API_KEY"))
	// helper.HandleError(viper.BindEnv("API_SECRET"))
	// helper.HandleError(viper.BindEnv("USERNAME"))
	// helper.HandleError(viper.BindEnv("PASSWORD"))
	//
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	// }

	// log settings
	internal.InitLogger()
}
