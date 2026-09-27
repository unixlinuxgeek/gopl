// Exercise 1.3:
//
// Experiment to measure the difference in running time between our
// potentially inefficient versions and the one that uses strings.Join.
// (Section 1.6 illustrates part of the time package,
// and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "argument count is 0\n")
		os.Exit(1)
	}

	out := Echo2()

	fmt.Println(out)
}

func Echo2() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}
