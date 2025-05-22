/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Taiwopeter-babs/clig/todo"
	"github.com/spf13/cobra"
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

	items, err := todo.ReadItems(datafile)

	if err != nil {
		fmt.Println(err)
	}

	for _, val := range args {
		// fmt.Println(val)
		item := todo.Item{Text: val}
		item.SetPriority(priority)
		items = append(items, item)
	}

	err = todo.SaveItems(datafile, items)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", items)
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
