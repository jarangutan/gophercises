/*
Copyright © 2024 Jose Aranguren
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
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
		fmt.Println("add called")
		s := strings.Join(args, " ")
		db, errDb := bolt.Open("my.db", 0600, nil)
		if errDb != nil {
			log.Fatal(errDb)
		}
		defer db.Close()

		errUpdate := db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
			if err != nil {
				log.Fatal(err)
				return err
			}

			task := Task{Task: s, Completed: false}
			taskJson, err := json.Marshal(task)
			if err != nil {
				log.Fatal(err)
				return err
			}

			t := time.Now().Format(time.RFC3339)
			return b.Put([]byte(t), []byte(taskJson))
		})
		if errUpdate != nil {
			log.Fatal(errUpdate)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
