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

	in, err := DetermineIn()
	if err != nil {
		log.Fatal(err)
	}

	defer in.Close()

	result, err := UniqNoArgs(in)
	if err != nil {
		log.Fatal(err)
	}

	out, err := DetermineOut()
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	for _, line := range result {
		out.Write([]byte(line))
	}
}

func DetermineIn() (io.ReadCloser, error) {
	flag.Parse()
	var in io.ReadCloser

	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("error opening file: %g", err)
		}

		in = f
	} else {
		in = os.Stdin
	}

	return in, nil
}

func DetermineOut() (io.WriteCloser, error) {
	flag.Parse()
	var out io.WriteCloser

	if filename := flag.Arg(1); filename != "" {
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return nil, fmt.Errorf("error opening file: %g", err)

		}
		out = f
	} else {
		out = os.Stdout
	}

	return out, nil
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
