/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/Taiwopeter-babs/clig/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: listRun,
	}

	doneOpt, allOpt bool
)

func listRun(cmd *cobra.Command, args []string) {
	var datafileName string = *todo.AllConstants.DataFileName

	items, err := todo.ReadItems(viper.GetString(datafileName))

	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(todo.ByPriority(items))

	writer := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	for _, item := range items {
		if allOpt || item.Done == doneOpt {
			fmt.Fprintln(writer, item.Label()+"\t"+item.PrettyDone()+"\t"+item.PrettyP()+"\t"+item.Text+"\t")
		}
	}

	writer.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos")

	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show All Todos")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
