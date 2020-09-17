// ex1.4 prints the file name as well as count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// holds lines and their counts
	counts := make(map[string]int)
	// format of line:[file1, file2]
	inFiles := make(map[string][]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, inFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dupmod: %v\n", err)
				continue
			}
			countLines(f, counts, inFiles)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, inFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, inFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()

		counts[line]++

		filename := f.Name()
		if !inSlice(filename, inFiles[line]) {
			inFiles[line] = append(inFiles[line], filename)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

func inSlice(filename string, files []string) bool {
	for _, f := range files {
		if f == filename {
			return true
		}
	}
	return false
}
