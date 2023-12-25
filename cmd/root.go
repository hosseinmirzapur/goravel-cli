/*
Copyright © 2023 Hossein Mirzapur hosseinmirzapur@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hosseinmirzapur/goravel-cli/cmd/artisan"
	"github.com/hosseinmirzapur/goravel-cli/cmd/start"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goravel-cli",
	Short: "Goravel CLI Application",
	Long:  "Start your amazing Goravel application in no time",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addCommands() {
	rootCmd.AddCommand(start.StartCmd)
	rootCmd.AddCommand(artisan.ArtisanCmd)
}

func init() {
	cobra.OnInitialize(initConfig)

	addCommands()

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goravel-cli.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".goravel-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".goravel-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
