/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"task-cli/services"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <ID> <YOUR TASK NAME>",
	Short: "Update your task based on its ID.",
	Long: `This command allows you to update a task using its ID. 
Example usage:
	task-cli update 1 "New task" 
	task-cli update 099 "99th task"
	`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("invalid arguments: expected <ID> <YOUR TASK NAME>")
		}
		return nil
	},
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("invalid ID: %v\n", err)
			return
		}
		description := args[1]
		if err := services.UpdateTask(id, description); err != nil {
			fmt.Printf("unable to update task: %v\n", err)
			return
		}
		fmt.Println("Task updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
