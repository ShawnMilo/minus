package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var errNoInput = errors.New("no input")

func isInteractive() (bool, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	return stat.Mode()&os.ModeCharDevice != 0, nil
}

func inputReader() (io.Reader, error) {
	if b, err := isInteractive(); err == nil {
		if !b {
			return os.Stdin, nil
		}
	}

	readers := make([]io.Reader, 0, 1)
	for _, filename := range os.Args[1:] {
		if fileExists(filename) {
			f, err := os.Open(filename)
			if err != nil {
				continue
			}
			readers = append(readers, f)
		}
	}
	if len(readers) == 0 {
		return io.MultiReader(readers...), errNoInput
	}
	return io.MultiReader(readers...), nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func main() {
	filter := os.Args[1:]
	fmt.Fprintf(os.Stderr, "filter: %v\n", filter)
	reader, err := inputReader()
	var total int
	var skipped int

	if err != nil {
		log.Fatalf("failed to read: %s", err)
	}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		txt := scanner.Text()
		var skip bool
		for _, f := range filter {
			if strings.Contains(txt, f) {
				skip = true
				skipped++
				break
			}
		}
		if skip {
			continue
		}
		total++
		fmt.Println(txt)
	}
	fmt.Fprintf(os.Stderr, "total: %d, skipped: %d, remaining: %d\n", total, skipped, total-skipped)
}
