package lines

import (
	"bufio"
	"fmt"
	"os"
)

var lines [][]byte
var watched = 0

func extractTail(s *bufio.Scanner, mux int) [][]byte {
	lines = nil
	for s.Scan() {
		watched++
		if mux == 0 {
			continue
		}
		if len(lines) > mux-1 {
			lines = append(lines[1:])
		}
		lines = append(lines, s.Bytes())
	}
	return lines
}

func extractChange(s *bufio.Scanner) [][]byte {
	lines = nil
	i := 0
	for s.Scan() {
		i++
		if watched < i {
			lines = append(lines, s.Bytes())
			watched++
		}
	}
	return lines
}

// Create create lines for output.
func Create(path string, n int) [][]byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(2)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// check first run
	if watched == 0 {
		return extractTail(scanner, n)
	}
	return extractChange(scanner)
}
