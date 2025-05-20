/*
Copyright © 2025 Taiwo Babalola
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/mitchellh/go-homedir"
)

var datafile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clig",
	Short: "Clig is a todo application",
	Long:  `Clig will help you get more done in less time`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set datafile using --datafile")
	}

	rootCmd.PersistentFlags().StringVar(
		&datafile,
		"datafile",
		home+string(os.PathSeparator)+".clig-todos.json",
		"datafile to store todos",
	)

}
