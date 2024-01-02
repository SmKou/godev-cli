package main

import (
	"bufio" // read text
	"flag"  // command line flags
	"fmt"   // print formatted output
	"io"    // io.Reader interface
	"os"    // os resources: Stdin
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("c", false, "Count bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	counter := 0
	for scanner.Scan() {
		counter++
	}
	return counter
}
