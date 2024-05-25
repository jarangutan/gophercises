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

var pendingCmd = &cobra.Command{
	Use:   "pending",
	Short: "List all of your pending tasks for today",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		minTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		maxTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 9999, now.Location())

		tasks, err := db.ListTasksWithinTimeRange(minTime, maxTime)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}

		pendingTasks := db.FilterTasks(tasks, db.FilterPending)
		if pendingTasks == nil {
			fmt.Println("No tasks are pending today.")
			return
		}

		fmt.Printf("You have finished the following tasks today:\n")
		for _, t := range pendingTasks {
			fmt.Printf("- %s\n", t.Task)
		}
	},
}

func init() {
	rootCmd.AddCommand(pendingCmd)
}
