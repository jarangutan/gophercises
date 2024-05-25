/*
Copyright © 2024 Jose Aranguren
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"4d63.com/homedir"
	"github.com/jarangutan/gophercises/task/cmd"
	"github.com/jarangutan/gophercises/task/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "task.db")
	must(db.Init(dbPath))
	cmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
