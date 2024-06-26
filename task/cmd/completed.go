/*
Copyright © 2024 Jose Aranguren
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/jarangutan/gophercises/task/db"
	"github.com/spf13/cobra"
)

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List all of your completed tasks for today",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		minTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		maxTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 9999, now.Location())

		tasks, err := db.ListTasksWithinTimeRange(minTime, maxTime)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}

		completedTasks := db.FilterTasks(tasks, db.FilterCompleted)
		if completedTasks == nil {
			fmt.Println("No tasks have been completed today.")
			return
		}

		fmt.Printf("You have finished the following tasks today:\n")
		for _, t := range completedTasks {
			fmt.Printf("- %s\n", t.Task)
		}
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)
}
