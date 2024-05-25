/*
Copyright © 2024 Jose Aranguren
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jarangutan/gophercises/task/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You havuwue no tasks yet")
			return
		}

		fmt.Printf("You have the following tasks:\n")
		for i, t := range tasks {
			if t.Completed {
				fmt.Printf("%d. %s (completed)\n", i+1, t.Task)
				continue
			}
			fmt.Printf("%d. %s\n", i+1, t.Task)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
