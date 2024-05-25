package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

// fff.jkl.gh

func caesarCipher(s string, k int32) string {
	var ret []rune
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			newCh := ch + k
			if newCh > 'z' {
				fmt.Println("lower char rotate", (((ch + k) % 'z') + ('a' - 1)))
				newCh = ((ch + k) % 'z') + ('a' - 1)
			}
			fmt.Println("lower char", string(ch))
			ret = append(ret, newCh)
		} else if ch >= 'A' && ch <= 'Z' {
			newCh := ch + k
			if newCh > 'Z' {
				fmt.Println("lower char rotate", (((ch + k) % 'Z') + ('A' - 1)))
				newCh = ((ch + k) % 'Z') + ('A' - 1)
			}
			fmt.Println("upper char", string(ch))
			ret = append(ret, newCh)
		} else {
			fmt.Println("any char", string(ch))
			ret = append(ret, ch)
		}
	}
	return string(ret)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	readLine(reader)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
