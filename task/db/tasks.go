package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"

	bolt "go.etcd.io/bbolt"
)

var taskBucket = []byte("Tasks")
var db *bolt.DB

type Task struct {
	Key       string
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		t := Task{Task: task, Completed: false}
		taskJson, err := json.Marshal(t)
		if err != nil {
			return err
		}

		key := time.Now().Format(time.RFC3339)
		return b.Put([]byte(key), []byte(taskJson))
	})
}

func UpdateTask(key []byte, task Task) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		taskJson, err := json.Marshal(task)
		if err != nil {
			return err
		}
		return b.Put([]byte(key), []byte(taskJson))
	})
}

func ListTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.ForEach(func(k, v []byte) error {
			var task Task
			err := json.Unmarshal(v, &task)
			if err != nil {
				return err
			}
			tasks = append(tasks, task)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func ListCompletedTasksWithinTimeRange(minTime time.Time, maxTime time.Time) ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("Tasks")).Cursor()
		min := []byte(minTime.Format(time.RFC3339))
		max := []byte(maxTime.Format(time.RFC3339))

		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			var task Task
			err := json.Unmarshal(v, &task)
			if err != nil {
				return err
			}
			if task.Completed {
				tasks = append(tasks, task)
			}
		}
		return nil
	})
	return tasks, err
}

func GetTask(key []byte) (Task, error) {
	task := Task{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		v := b.Get(key)
		if v == nil {
			return errors.New("No record found with given key")
		}
		err := json.Unmarshal(v, &task)
		if err != nil {
			return err
		}
		return nil
	})
	return task, err
}

func GetKeyByIndex(index int) ([]byte, error) {
	var key []byte
	i := 0

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			if i == index-1 {
				key = k
				break
			}
			i++
		}
		return nil
	})
	return key, err
}

func DeleteTask(key []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(key)
	})
}
