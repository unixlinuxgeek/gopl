// Exercise 1.2:
// Modify the echo program to print the index and value of each of its arguments, one per line.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		if i == len(os.Args)-1 {
			sep = ""
		} else {
			sep = "\n"
		}
		s += strconv.Itoa(i) + "." + os.Args[i] + sep
	}
	fmt.Println(s)
}
