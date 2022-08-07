package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	result, err := UniqNoArgs(os.Stdin)

	if err != nil {
		log.Fatal(err)
	}

	for _, line := range result {
		fmt.Print(line)
	}
}

func UniqNoArgs(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	last := ""
	var result []string

	for scanner.Scan() {
		scannedLine := scanner.Text()
		if scannedLine != last {
			result = append(result, scannedLine+"\n")
			last = scannedLine
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading standard input: %g", err)
	}

	return result, nil

}
