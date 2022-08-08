package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var in io.Reader

	flag.Parse()
	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer f.Close()

		in = f
	} else {
		in = os.Stdin
	}

	result, err := UniqNoArgs(in)
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
