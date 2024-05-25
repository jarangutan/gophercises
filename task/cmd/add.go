/*
Copyright © 2024 Jose Aranguren
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/jarangutan/gophercises/task/db"
	"github.com/spf13/cobra"
)

type Task struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		err := db.CreateTask(task)

		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}

		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
