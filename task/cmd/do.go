/*
Copyright © 2024 Jose Aranguren
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"4d63.com/homedir"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, errConv := strconv.Atoi(args[0])
		if errConv != nil {
			log.Fatal(errConv)
		}

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

		var key []byte
		var task Task
		i := 0
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Tasks"))
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				if i == index-1 {
					key = k
					err := json.Unmarshal(v, &task)
					if err != nil {
						log.Fatal("oops", err)
						return err
					}
					break
				}
			}
			return nil
		})

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Tasks"))
			task.Completed = true
			taskJson, errMarshal := json.Marshal(task)
			if errMarshal != nil {
				log.Fatal(errMarshal)
				return errMarshal
			}

			err := b.Put([]byte(key), []byte(taskJson))
			return err
		})
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
