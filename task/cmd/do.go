/*
Copyright © 2024 Jose Aranguren
*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jarangutan/gophercises/task/db"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:     "do [space separated indexes]",
	Short:   "Mark a task on your TODO list as complete",
	Example: "do 1 2 3",
	RunE: func(cmd *cobra.Command, args []string) error {
		var indexes []int
		for _, a := range args {
			index, err := strconv.Atoi(a)
			if err != nil {
				return errors.New(fmt.Sprintf("Failed to parse argument. Argument \"%v\" is not a number.\n", a))
			}
			indexes = append(indexes, index)
		}

		for _, i := range indexes {
			key, err := db.GetKeyByIndex(i)
			if err != nil {
				fmt.Println("Something went wrong:", err)
				return err
			}
			if key == nil {
				fmt.Printf("No task found with an index of %d.\n", i)
				return nil
			}

			task, err := db.GetTask(key)
			if err != nil {
				fmt.Println("Something went wrong:", err)
				return err
			}

			task.Completed = true

			err = db.UpdateTask(key, task)
			if err != nil {
				fmt.Println("Something went wrong updating task:", err)
				return err
			}
			fmt.Printf("Completed task \"%s\".\n", task.Task)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
