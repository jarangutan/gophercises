package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := flag.String("input", "problems.csv", "path to problems csv")
	flag.Parse()

	fmt.Println("word: ", *input)

	file, err := os.Open(*input)
	if err != nil {
		fmt.Printf("File was not opened due to: %v", err)
	}

	parser := csv.NewReader(file)
	scanner := bufio.NewScanner(os.Stdin)

	correct := 0
	questions := 0

	for {
		record, err := parser.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("Ran into a problem reading line from file: %v", err)
		}

		questions += 1

		fmt.Println(record[0])
		scanner.Scan()
		in := scanner.Text()

		input := strings.TrimSpace(in)

		if len(input) == 0 {
			fmt.Println("Quitting")
			break
		}

		if input == record[1] {
			correct += 1
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	fmt.Printf("You got %v right out of %v", correct, questions)
}
