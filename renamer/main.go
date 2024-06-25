package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// mask legend
var x rune = 'X'
var n rune = 'N'

func match(filename string, mask string) (string, error) {
	if len(filename) != len(mask) {
		return "", fmt.Errorf("filename '%s' length did not match mask", filename)
	}
	var ext string
	var part []rune
	var number []rune

	pieces := strings.Split(filename, ".")
	ext = pieces[len(pieces)-1]

	fn := []rune(filename)

	for i, r := range mask {
		switch r {
		case x:
			part = append(part, fn[i])
		case n:
			number = append(number, fn[i])
		}
	}

	n, err := strconv.Atoi(string(number))
	if err != nil {
		return "", fmt.Errorf("'%s' didn't match mask", filename)
	}

	return fmt.Sprintf("%s - %d of %d.%s", string(part), n, 4, ext), nil
}

func main() {
	var dry bool
	flag.BoolVar(&dry, "dry", false, "whether or not this should be a real or dry run")
	dirname := flag.String("dir", ".", "JSON file with cyoa story")
	flag.Parse()
	fmt.Printf("Directory '%s'\n", *dirname)

	mask := "XXXXXXXX_NNN.txt"

	err := filepath.Walk(*dirname, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		newName, errMatch := match(info.Name(), mask)
		if errMatch != nil {
			return nil
		}
		dir := filepath.Dir(path)
		newPath := filepath.Join(dir, newName)
		fmt.Printf("mv %s => %s\n", path, newPath)
		if !dry {
			os.Rename(path, newPath)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Something went wrong")
		os.Exit(1)
	}

	fmt.Println("Done :-)")
}
