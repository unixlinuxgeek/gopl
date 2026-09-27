// Exercise 1.4:
// Modify dup2 to print the names of all files in which each duplicated line occurs.
/*
Run:

	$ go run dup2.go testdata/file1.txt  testdata/file2.txt
	4       testdata/file1.txt/Melisa Patel
	2       testdata/file1.txt/Isabel Porter
	2       testdata/file1.txt/Maja Welsh
	3       testdata/file2.txt/Zunaisha Webb

*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[f.Name()+"/"+input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
