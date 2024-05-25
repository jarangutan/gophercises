/*
Copyright © 2024 Jose Aranguren
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/jarangutan/gophercises/task/db"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("%v is not a number.\n", args[0])
		}

		key, err := db.GetKeyByIndex(index)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		if key == nil {
			fmt.Printf("No task found with an index of %d.\n", index)
			return
		}

		task, err := db.GetTask(key)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}

		task.Completed = true

		err = db.UpdateTask(key, task)
		if err != nil {
			fmt.Println("Something went wrong updating task:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
