// Exercise 1.2:
// Modify the echo program to print the index and value of each of its arguments, one per line.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		if i > 0 {
			sep = "\n"
		}
		s += sep + strconv.Itoa(i+1) + "." + arg
	}
	fmt.Println(s)
}
