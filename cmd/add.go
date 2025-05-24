/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Taiwopeter-babs/clig/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new item and add to the list`,
	Run:   addRun,
}

// Runs the add command
func addRun(cmd *cobra.Command, args []string) {
	var err error

	var configDataFileName = *todo.AllConstants.ConfigDataFileName

	items, _ := todo.ReadItems(viper.GetString(configDataFileName))

	for _, val := range args {
		item := todo.Item{Text: val}
		item.SetPriority(priority)
		items = append(items, item)
	}

	err = todo.SaveItems(viper.GetString(configDataFileName), items)

	if err != nil {
		return
	}

	fmt.Printf("%#v\n", items)
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")

}
