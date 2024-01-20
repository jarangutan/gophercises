package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
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

	records, err := parser.ReadAll()

	if err != nil {
		fmt.Printf("Ran into a problem reading line from file: %v", err)
	}

	correct := 0

	for _, record := range records {
		question, answer := record[0], record[1]

		fmt.Println(question)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if len(input) == 0 {
			fmt.Println("Quitting")
			break
		}

		if input == answer {
			correct += 1
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	fmt.Printf("You got %v right out of %v", correct, len(records))
}
