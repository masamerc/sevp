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
	Version: "0.1.0",
	Short:   "sevp: pick and switch environement variables.",
	Long:    `sevp: pick and switch environement variables.`,
	Args: cobra.MatchAll(
		cobra.MinimumNArgs(0),
		cobra.MaximumNArgs(1),
	),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			selectorChoice := args[0]
			// config parse
			selectors := internal.GetSelectors()

			// check if the selector exists in the map keys
			if selector, ok := selectors[selectorChoice]; ok {
				fmt.Println(selector.Name)
				fmt.Println(selector.ReadConfig)
				fmt.Println(selector.TargetVar)
				fmt.Println(selector.PossibleValues)

				app := app.NewApp(selector.PossibleValues, selector.TargetVar)

				err := app.Run()

				internal.FailOnError(
					"Error running app",
					err,
				)
			} else {
				fmt.Println("Selector not found in config")
			}
		} else {
			// default route
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
			app := app.NewApp(profiles, "AWS_PROFILE")

			err = app.Run()
			internal.FailOnError(
				"Error running app",
				err,
			)
		}
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
	// read in config
	internal.ParseConfig()

	// log settings
	internal.InitLogger()
}
