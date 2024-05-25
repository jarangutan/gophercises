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

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a task on your TODO list",
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

		var keys [][]byte
		for _, i := range indexes {
			key, err := db.GetKeyByIndex(i)
			if err != nil {
				fmt.Println("Something went wrong:", err)
				return
			}
			if key == nil {
				fmt.Printf("No task found with an index of %d.\n", i)
				continue
			}
			keys = append(keys, key)
		}

		for _, k := range keys {
			task, err := db.GetTask(k)
			if err != nil {
				fmt.Println("Something went wrong:", err)
				return
			}

			err = db.DeleteTask(k)
			if err != nil {
				fmt.Println("Something went wrong:", err)
			}

			fmt.Printf("You have deleted task \"%s\".\n", task.Task)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
