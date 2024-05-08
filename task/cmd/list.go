/*
Copyright © 2024 Jose Aranguren
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		db, errDb := bolt.Open("my.db", 0600, nil)
		if errDb != nil {
			log.Fatal(errDb)
		}
		defer db.Close()

		var tasks []Task
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Tasks"))
			b.ForEach(func(k, v []byte) error {
				var task Task
				err := json.Unmarshal(v, &task)
				if err != nil {
					log.Fatal("oops", err)
					return err
				}
				tasks = append(tasks, task)
				return nil
			})
			return nil
		})

		fmt.Printf("You have the following tasks:\n")
		for i, v := range tasks {
			fmt.Printf("%d. %s\n", i+1, v.Task)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
