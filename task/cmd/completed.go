/*
Copyright © 2024 Jose Aranguren
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"4d63.com/homedir"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List all of your completed tasks for today",
	Run: func(cmd *cobra.Command, args []string) {
		homepath, errHomedir := homedir.Dir()
		if errHomedir != nil {
			panic("Home dir not found!")
		}
		dbpath := fmt.Sprintf("%s/task.db", homepath)
		db, errDb := bolt.Open(dbpath, 0600, nil)
		if errDb != nil {
			log.Fatal(errDb)
		}
		defer db.Close()

		now := time.Now()
		minTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Format(time.RFC3339)
		maxTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 9999, now.Location()).Format(time.RFC3339)

		db.View(func(tx *bolt.Tx) error {
			c := tx.Bucket([]byte("Tasks")).Cursor()
			min := []byte(minTime)
			max := []byte(maxTime)

			fmt.Printf("You have finished the following tasks today:\n")
			for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
				var task Task
				err := json.Unmarshal(v, &task)
				if err != nil {
					log.Fatal("oops", err)
					return err
				}
				if task.Completed {
					fmt.Printf("- %s\n", task.Task)
				}
			}
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
