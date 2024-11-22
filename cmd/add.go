/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"task-cli/services"

	"github.com/spf13/cobra"
)

var Task string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <YOUR_TASK_NAME>",
	Short: "It will add a new task",
	Long: `This command will help to add a new task
	Eg:-
	task-cli add "Buy groceries"
	task-cli add "Write a poem"
	`,
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("exactly 1 argument is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id, err := services.AddTask(args[0])
		if err != nil {
			fmt.Printf("Error: unable to add task: %v\n", err)
			return
		}

		fmt.Printf("Task added successfully (ID:%v)\n", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().StringVarP(&Task, "task", "t", "", "Add your task name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
