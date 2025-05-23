/*
Copyright © 2025 Taiwo Babalola
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Taiwopeter-babs/clig/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configName string = ".clig"
	configPath string = "$HOME"
)

var (

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "clig",
		Short: "Clig is a todo application",
		Long:  `Clig will help you get more done in less time`,
	}

	datafile string
	cfgFile  string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using Config file: ", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(
		&datafile,
		"datafile",
		*todo.AllConstants.Filename,
		"datafile to store todos",
	)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.clig.yaml)")

}
