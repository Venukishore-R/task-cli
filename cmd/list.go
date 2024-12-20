/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"task-cli/services"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all the tasks and list the tasks based on the status provided`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("invalid arguments: expected only one argument")
		}
		return nil
	},
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		status := "all"
		if len(args) == 1 {
			status = args[0]
		}

		tasks, err := services.ListTasks(status)
		if err != nil {
			fmt.Printf("unable to list task: %v\n", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

		// Print header
		fmt.Fprintln(writer, "ID\tDescription\t\t\tStatus")
		fmt.Fprintln(writer, "----\t---------------------------------\t--------")

		// Print rows
		for _, task := range tasks {
			fmt.Fprintf(writer, "%d\t%s\t\t\t%s\n", task.Id, task.Description, task.Status)
		}

		// Flush the writer to output the formatted table
		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
