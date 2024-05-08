/*
Copyright © 2024 Jose Aranguren
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do called")
		index, errConv := strconv.Atoi(args[0])
		if errConv != nil {
			log.Fatal(errConv)
		}

		db, errDb := bolt.Open("my.db", 0600, nil)
		if errDb != nil {
			log.Fatal(errDb)
		}
		defer db.Close()

		var key []byte
		var task Task
		i := 0
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Tasks"))
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				if i == index-1 {
					fmt.Printf("key=%s, value=%s\n", k, v)
					key = k
					err := json.Unmarshal(v, &task)
					if err != nil {
						log.Fatal("oops", err)
						return err
					}
					fmt.Printf("key=%s, value=%s\n", key, v)
					break
				}
			}
			return nil
		})

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Tasks"))
			err := b.Delete(key)
			return err
		})
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
