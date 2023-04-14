package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func countBytes(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanBytes)
	return countScan(*scanner)
}

func countLines(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	return countScan(*scanner)
}

func countWords(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	return countScan(*scanner)
}

func countScan(scanner bufio.Scanner) int {
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func main() {
	bytes := flag.Bool("b", false, "count bytes")
	lines := flag.Bool("l", false, "count lines")
	input := os.Stdin
	count := 0

	flag.Parse()

	if *bytes {
		count = countBytes(input)
	} else if *lines {
		count = countLines(input)
	} else {
		count = countWords(input)
	}

	fmt.Println(count)
}
