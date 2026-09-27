// Exercise 1.2:
// Modify the echo program to print the index and value of each of its arguments, one per line.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for i := 0; i != len(os.Args); i++ {
		os.Args[i] = strconv.Itoa(i) + "." + os.Args[i]
	}

	fmt.Println(strings.Join(os.Args[1:], "\n"))
}
