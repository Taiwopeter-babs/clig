/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Taiwopeter-babs/clig/todo"
	"github.com/spf13/cobra"
)

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
		log.Fatal(err)
	}

	for _, val := range args {
		fmt.Println(val)
		items = append(items, todo.Item{Text: val})
	}

	err = todo.SaveItems(datafile, items)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", items)
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
