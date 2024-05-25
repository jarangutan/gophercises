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
	Run: func(cmd *cobra.Command, args []string) {
		var indexes []int
		for _, a := range args {
			index, err := strconv.Atoi(a)
			if err != nil {
				fmt.Printf("Failed to parse argument. Argument \"%v\" is not a number.\n", a)
				return
			}
			indexes = append(indexes, index)
		}

		for _, i := range indexes {
			key, err := db.GetKeyByIndex(i)
			if err != nil {
				fmt.Println("Something went wrong:", err)
				return
			}
			if key == nil {
				fmt.Printf("No task found with an index of %d.\n", i)
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
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
