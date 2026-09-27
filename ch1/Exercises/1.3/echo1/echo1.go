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

	out := Echo1()

	fmt.Println(out)
}

func Echo1() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}
