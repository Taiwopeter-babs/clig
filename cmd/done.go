/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/Taiwopeter-babs/clig/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"done"},
	Short:   "Mark Item as Done",
	Run:     doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {

	items, err := todo.ReadItems(datafile)

	if err != nil {
		log.Fatalln(err)
	}

	if len(args) < 1 {
		log.Fatalln("One argument must be provided is not a valid label\n", err)
	}

	arg, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	if arg > 0 && arg < len(items) {
		items[arg-1].Done = true
		fmt.Printf("%q %v\n", items[arg-1].Text, "marked done!")

		sort.Sort(todo.ByPriority(items))
		todo.SaveItems(datafile, items)
	} else {
		log.Println("(", arg, ")", "does not match any item", "|| Amount of items =", len(items))
	}

}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
