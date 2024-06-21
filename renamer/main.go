package main

import (
	"fmt"
	"strconv"
	"strings"
)

// mask
var x rune = 'X'
var n rune = 'N'

func parseFileName(filename string, mask string) (string, error) {
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
		return "", err
	}

	return fmt.Sprintf("%s - %d of %d.%s", string(part), n, 4, ext), nil
}

func main() {
	filename := "birthday_001.txt"
	mask := "XXXXXXXX_NNN.txt"
	result, err := parseFileName(filename, mask)
	if err != nil {
		panic(err)
	}
	println(result)
}
